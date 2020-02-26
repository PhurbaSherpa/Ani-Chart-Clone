package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=phrbshrp password=Nepal123 dbname=anichart sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	table, err := db.Prepare(
		`CREATE TABLE Anime (
							id 		SERIAL 		PRIMARY KEY		 NOT NULL,
							title 					TEXT 					 NOT NULL,
							description			TEXT					 NOT NULL,
							date 						VARCHAR(200)	 NOT NULL,
							season 					TEXT  				 NOT NULL
		)`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = table.Exec()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Anime table created successfully")
	}

	table, err = db.Prepare(
		`CREATE TABLE Genre (
							id 		SERIAL 		PRIMARY KEY		 NOT NULL,
							type 					  TEXT 					 NOT NULL
		)`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = table.Exec()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Genre table created successfully")
	}

	table, err = db.Prepare(
		`CREATE TABLE AnimeGenre (
								id 		SERIAL 		PRIMARY KEY		 NOT NULL,
								animeId  				INT,
								genreId					INT
			)`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = table.Exec()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("AnimeGenre table created successfully")
	}

	defer db.Close()

}
