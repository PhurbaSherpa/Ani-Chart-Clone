package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type anime struct {
	title       string
	description string
	date        string
	season      string
}

type genre struct {
	genre string
}

type animegenre struct {
	animeid int
	genreid int
}

func main() {
	connStr := "user=phrbshrp password=Nepal123 dbname=anichart sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	animes := []anime{
		{"Haikyuu!! To the TOP!!", "The fourth season of Haikyuu!! The Karasuno High School Volleyball Club finally won their way into the nationals after an intense battle for the Miyagi Prefecture Spring Tournament qualifiers. As they were preparing for the nationals, Kageyama is invited to go to All-Japan Youth Training Camp. At the same time, Tsukishima is invited to go to a special rookie select training camp for first-years in Miyagi Prefecture. Hinata feels panic that he’s being left behind.", "Jan 11, 2020", "Winter"}, {"DR.STONE", "After five years of harboring unspoken feelings, highschooler Taiju Ooki is finally ready to confess his love to Yuzuriha Ogawa. Just when Taiju begins his confession however, a blinding green light strikes the Earth and petrifies mankind around the world turning every single human into stone. Several millennia later, Taiju awakens to find the modern world completely nonexistent, as nature has flourished in the years humanity stood still. Among a stone world of statues, Taiju encounters one other living human: his science loving friend Senkuu, who has been active for a few months. Taiju learns that Senkuu has developed a grand scheme to launch the complete revival of civilization with science.", "Jul 5, 2019", "Summer"}, {"Boku no Hero Academia 4", "The villain world teeters on the brink of war now that All For One is out of the picture. Shigaraki of the League of Villains squares off with Overhaul of the yakuza, vying for total control of the shadows. Meanwhile, Deku gets tangled in another dangerous internship as he struggles to keep pace with his upperclassmanMirio.", "Oct 12, 2019", "Fall"}, {"Kimetsu no Yaiba", "Ever since the death of his father, the burden of supporting the family has fallen upon Tanjirou Kamado. Though living impoverished on a remote mountain, the Kamado family are able to enjoy a relatively peaceful life. One day, Tanjirou arrives back home and his whole family has been slaughtered. Worse still, the sole survivor is his sister Nezuko, has been turned into a bloodthirsty demon. Consumed by rage and hatred, Tanjirou swears to avenge his family and stay by his only remaining sibling.", "Apr 6, 2019", "Spring"},
	}

	genres := []genre{
		{"Action"}, {"Comedy"}, {"Fantasy"}, {"Adventure"}, {"Drama"}, {"Sports"}, {"Supernatural"}, {"Mysetry"}, {"Love"}, {"Sci-fi"}}

	animegenres := []animegenre{
		{1, 1}, {1, 2}, {1, 5}, {1, 6}, {2, 1}, {2, 2}, {2, 4}, {2, 10}, {3, 1}, {3, 2}, {3, 10}, {4, 1}, {4, 3}, {4, 4}, {4, 5}, {4, 7}, {4, 8}, {4, 10},
	}

	for _, anime := range animes {
		sqlStaement := fmt.Sprintf(`
		INSERT INTO anime (title, description, date, season)
		VALUES ('%s','%s','%s','%s')`, anime.title, anime.description, anime.date, anime.season)
		_, err = db.Exec(sqlStaement)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, genre := range genres {
		sqlStaement := fmt.Sprintf(`
		INSERT INTO genre (genre) VALUES ('%s')`, genre.genre)
		_, err = db.Exec(sqlStaement)
		if err != nil {
			log.Fatalln(err)
		}
	}

	for _, animegenre := range animegenres {
		sqlStaement := fmt.Sprintf(`
		INSERT INTO animegenre (animeid,genreid) VALUES ('%d', '%d')`, animegenre.animeid, animegenre.genreid)
		_, err = db.Exec(sqlStaement)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
