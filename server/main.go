package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var connStr string = "user=phrbshrp password=Nepal123 dbname=anichart sslmode=disable"

type Anime struct {
	Id          int
	Title       string
	Description string
	Date        string
	Season      string
	Imageurl    string
}

type Character struct {
	Id          int
	AnimeId     int
	Name        string
	Imageurl    string
	Role        string
	Description string
}

func handleCharactersRoute(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	switch req.Method {
	case "GET":
		var characters []*Character
		vars := mux.Vars(req)
		uri := vars["id"]
		sqlStatement := `SELECT * FROM character
												WHERE character.animeId = $1`
		rows, err := db.Query(sqlStatement, uri)
		defer rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			a := new(Character)
			err := rows.Scan(&a.Id, &a.AnimeId, &a.Name, &a.Imageurl, &a.Role, &a.Description)
			if err != nil {
				log.Fatalln(err)
			}
			characters = append(characters, a)
		}

		fmt.Println("200 SUCCESS", characters)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(characters)
		if err != nil {
			log.Fatalln(err)
		}
		return
	case "POST":
		fmt.Fprint(w, "POST")
	default:
		log.Fatalln("Bad Request")
	}
}

func handleGenreRoute(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	switch req.Method {
	case "GET":
		var genres []string
		vars := mux.Vars(req)
		uri := vars["id"]

		sqlStatement := `SELECT genre
											FROM animegenre
											INNER JOIN genre
											ON animegenre.genreId = genre.id
											INNER JOIN
												anime
												ON animegenre.animeId = anime.id
												WHERE anime.id = $1
											`
		rows, err := db.Query(sqlStatement, uri)
		defer rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			var genre string

			err := rows.Scan(&genre)

			if err != nil {
				log.Fatalln(err)
			}
			genres = append(genres, genre)
		}
		fmt.Println("200 SUCCESS")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(genres)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
}

func handleAnimeRoute(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	switch req.Method {
	case "GET":
		var animes []*Anime
		sqlStatement := `SELECT * FROM anime`
		rows, err := db.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			a := new(Anime)
			err := rows.Scan(&a.Id, &a.Title, &a.Description, &a.Date, &a.Season, &a.Imageurl)
			if err != nil {
				log.Fatalln(err)
			}
			animes = append(animes, a)
		}
		fmt.Println("200 SUCCESS")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(animes)
		if err != nil {
			log.Fatalln(err)
		}
		return
	case "POST":
		fmt.Fprint(w, "POST")
	default:
		log.Fatalln("Bad Request")
	}
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/animes", handleAnimeRoute)
	router.HandleFunc("/api/genres/{id:.+}", handleGenreRoute)
	router.HandleFunc("/api/characters/{id:.+}", handleCharactersRoute)

	spa := spaHandler{staticPath: "./static", indexPath: "./public/index.html"}
	router.PathPrefix("/animes/").Handler(spa)
	router.PathPrefix("/genres/").Handler(spa)
	router.PathPrefix("/characters/").Handler(spa)
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
