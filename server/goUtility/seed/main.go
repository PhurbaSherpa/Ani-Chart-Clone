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
	imageurl    string
}

type genre struct {
	genre string
}

type animegenre struct {
	animeid int
	genreid int
}

type character struct {
	animeid     int
	name        string
	imageurl    string
	role        string
	description string
}

func main() {
	connStr := "user=phrbshrp password=Nepal123 dbname=anichart sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	animes := []anime{
		{"Haikyuu!! To the TOP!!", "The fourth season of Haikyuu!! The Karasuno High School Volleyball Club finally won their way into the nationals after an intense battle for the Miyagi Prefecture Spring Tournament qualifiers. As they were preparing for the nationals, Kageyama is invited to go to All-Japan Youth Training Camp. At the same time, Tsukishima is invited to go to a special rookie select training camp for first-years in Miyagi Prefecture. Hinata feels panic that he’s being left behind.", "Jan 11, 2020", "Winter", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx106625-UR22wB2NuNVi.png"}, {"DR.STONE", "After five years of harboring unspoken feelings, highschooler Taiju Ooki is finally ready to confess his love to Yuzuriha Ogawa. Just when Taiju begins his confession however, a blinding green light strikes the Earth and petrifies mankind around the world turning every single human into stone. Several millennia later, Taiju awakens to find the modern world completely nonexistent, as nature has flourished in the years humanity stood still. Among a stone world of statues, Taiju encounters one other living human: his science loving friend Senkuu, who has been active for a few months. Taiju learns that Senkuu has developed a grand scheme to launch the complete revival of civilization with science.", "Jul 5, 2019", "Summer", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx105333-5p1MKBlGxZFF.jpg"}, {"Boku no Hero Academia 4", "The villain world teeters on the brink of war now that All For One is out of the picture. Shigaraki of the League of Villains squares off with Overhaul of the yakuza, vying for total control of the shadows. Meanwhile, Deku gets tangled in another dangerous internship as he struggles to keep pace with his upperclassman Mirio.", "Oct 12, 2019", "Fall", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx104276-DplpGzgCoRZX.jpg"}, {"Kimetsu no Yaiba", "Ever since the death of his father, the burden of supporting the family has fallen upon Tanjirou Kamado. Though living impoverished on a remote mountain, the Kamado family are able to enjoy a relatively peaceful life. One day, Tanjirou arrives back home and his whole family has been slaughtered. Worse still, the sole survivor is his sister Nezuko, has been turned into a bloodthirsty demon. Consumed by rage and hatred, Tanjirou swears to avenge his family and stay by his only remaining sibling.", "Apr 6, 2019", "Spring", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx101922-PEn1CTc93blC.jpg"}, {"Darwins Game", "Kaname Sudo, an ordinary high school student, receives an invitation email to try a mysterious app called Darwins Game. Kaname, upon launching the app, is drawn into a game where players fight one another using superpowers called Sigils.", "Jan 4, 2020", "Winter", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx105190-1OG3lBqoIQ7o.png"}, {"Sword Art Online: Alicization - War of Underworld", "Despite their recent victory against Quinella, Kirito and Alice find themselves in another dangerous situation. The Ocean Turtle, the complex in which resides the machinery that sustains Underworld, is under attack by an unknown enemy, leading to Kirito being thrown into a comatose state. At the same time, in Underworld, forces from the Dark Territory are starting to move to invade the Human Empire.", "Oct 13, 2019", "Fall", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx108759-jcXbDf9BJTcb.jpg"}, {"Nanatsu no Taizai: Kamigami no Gekirin", "The third season of Nanatsu no Taizai.", "Oct 9, 2019", "Fall", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx108928-Ci5iFvCfWgVY.jpg"}, {"Plunderer", "Every human inhabiting the world of Alcia is branded by a “Count” or a number written on their body. Depending on how each person lives their life, this Count either goes up or down. For Hinas mother, her total drops to 0 and she’s pulled into the Abyss, never to be seen again. But her mothers last words send Hina on a quest to find a legendary hero from the Waste War—the fabled Ace!", "Jan 9, 2020", "Winter", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx101168-7bmHkVlFxBCS.png"}, {"Diamond no Ace Act II", "Picking up the next year after the end of the fall tournament, Seidou High School baseball team battle it out with new and old faces as they begin their tournament run at Koshien. (Source: crunchyroll)", "Apr 2, 2019", "Spring", "https://s4.anilist.co/file/anilistcdn/media/anime/cover/large/bx18689-HrvWZjpWvxMr.jpg"},
	}

	genres := []string{
		"Action", "Comedy", "Fantasy", "Adventure", "Drama", "Sports", "Supernatural", "Mysetry", "Love", "Sci-fi"}

	animegenres := []animegenre{
		{1, 1}, {1, 2}, {1, 5}, {1, 6}, {2, 1}, {2, 2}, {2, 4}, {2, 10}, {3, 1}, {3, 2}, {3, 10}, {4, 1}, {4, 3}, {4, 4}, {4, 5}, {4, 7}, {4, 8}, {4, 10}, {5, 1}, {5, 8}, {6, 1}, {6, 4}, {7, 1}, {7, 4}, {8, 1}, {8, 4}, {9, 5}, {9, 6},
	}

	characters := []character{{1, "Shouyou Hinata", "https://s4.anilist.co/file/anilistcdn/character/large/b64769-rcDs153W4wLg.png", "Main", "Hinata has an unusually high ability to jump. Since he was young, Hinata trained his legs to compensate for his lack of height. He joins Karusuno High School Volleyball team and amazes the team with his natural affinity as a spiker. Although he is a complete beginner and never had any experience playing in a volleyball team, he grows rapidly through each match."}, {1, "Tobio Kageyama", "https://s4.anilist.co/file/anilistcdn/character/large/b64771-cS4Bbbs8Yvpz.png", "Main", "A famous setter in his middle school who joins the Karusuno High Volleyball team. He was known as the King of the Court in his middle school because of his arrogant demeanor and the way he would order around his teammates. Although he still carries an overbearing attitude, he begins to change when he starts to play with Hinata Shouyou who views him as a rival."}, {1, "Kei Tsukishima", "https://s4.anilist.co/file/anilistcdn/character/large/b67329-BaxY9PNMEWjL.jpg", "Supporting", "A member of the Karasuno High volleyball club. He often looks down on Hinata stating that he should not only rely on his jump. He also teases Kageyama calling him The king which angers him to the core."}, {1, "Daichi Sawamura", "https://s4.anilist.co/file/anilistcdn/character/large/b67323-yFbLDAfJ5ysG.png", "Supporting", "The captain of the Karasuno High volleyball club."}, {2, "Senkuu Ishigami", "https://s4.anilist.co/file/anilistcdn/character/large/b124142-AoRCosqVNunw.png", "Main", "Senku is a genius with a vast amount of scientific knowledge. He is able to invent various kinds of tools and gadgets in a short amount of time with ease. As a child, he designed and built a functioning miniature rocket ship. Senku also possesses an analytical mind, being able to correctly discern the situation he is in. Even in his pre-teens, he was intelligent and sharp-minded enough to create advanced machines. Senkus physical abilities are very low and he would lose against someone like Kohaku in a brawl, something he himself admits."}, {2, "Kohaku", "https://s4.anilist.co/file/anilistcdn/character/large/b124143-DABUFs3cPNrP.png", "Main", "Kohaku appears to be well known in her village due to her combative abilities, as both Kinro and Ginro were afraid of her. She is very fast and agile, being able to attack with her knives and legs swiftly. Her strength and endurance are also remarkable, being able to carry a heavy pot filled with hot water on a daily basis. She is also fast enough to dodge a bolt of lightning. Chrome referenced her as one of the few members in her village with superhuman strength, while also being very nimble and agile. She is incredible proficient in close-quarters combat and was able to easily stop an attack from Kinro and Ginro with her quick reflexes. Kohaku also possesses incredible eyesight, being able to spot Tsukasa and Senku from afar."}, {2, "Tsukasa Shishiou", "https://s4.anilist.co/file/anilistcdn/character/large/b124144-8lTkTDoPePiG.png", "Mian", "He is essentially violent, but extremely careful and aware of his actions. He seems to not have any friends and does not care for any other human besides himself. More importantly, he tends to hate humans who control and manipulate the others. By his actions, it is possible to assume that Tsukasa hates any hierarchy or practice of power. His goal is to create a society without modern corrupting technological advancements and he uses any means to achieve that objective; creating his own empire modeled on his twisted views of humanity."}, {3, "Izuku Midoriya", "https://s4.anilist.co/file/anilistcdn/character/large/89028-CmMSkY9TcWZq.jpg", "Main", "Izuku is the main protagonist of Boku no Hero Academia. Though originally born without a Quirk, he manages to catch the attention of the legendary hero All Might, and has since become his close pupil and a student at Yuuei. He is a very timid person who gets scared a lot, although in recent events, he has shown to raise his voice when the heat is on. He is very diligent, strong willed and a bit of a dreamer, and because he loves Heroes, he tries to know everything about them and writes down in his notebooks all he knows about heroes, including other Yuuei students."}, {3, "Katsuki Bakugou", "https://s4.anilist.co/file/anilistcdn/character/large/n88892-SEZsi3Cc6ikC.png", "Main", "Katsuki is a very aggressive person​ and is Midoriya Izukus rival character. He has a superiority complex and wants to be first, to be the best at everything. He doesnt like when people look down on him as if they were superior to him, and he gets mad easily. He is extremely arrogant and cocky. He also thinks of people such as Midoriya Izuku as an obstacle on his path to becoming the greatest hero."}, {3, "Toshinori Yagi", "https://s4.anilist.co/file/anilistcdn/character/large/b89224-10gTIqVOrPun.jpg", "Supporting", "All Might is the No. 1 Hero and the Symbol of Peace. In his hero form, All Might is a very tall and over-muscular man whose design resembles the american hero comics. His hair spikes on top of his head and his whole body has more shadows than a normal human being. However, in his true form, he is an over-skinny man who spits a lot of blood out of his mouth and his clothes appear wider."}, {4, "Tanjirou Kamado", "https://s4.anilist.co/file/anilistcdn/character/large/b126071-DEPUfKs1oWsp.png", "Main", "Tanjirou is kind by nature and has been described by others as having very gentle eyes. He has a great deal of determination and will not give up once he has a goal to achieve, an example being finding a cure for Nezuko. Even though he is relatively strong on his own, he isnt opposed to asking others for help when he needs it. He is very protective of his friends and even more so of his younger sister. His biggest attribute is his ability to empathize with anyone, even demons. This often gives him hesitation when killing demons."}, {4, "Nezuko Kamado", "https://s4.anilist.co/file/anilistcdn/character/large/b127518-7b4XKjasQE29.png", "Main", "Transformed into a demon and left mute, Nezuko travels with her older brother Tanjirou in a box on his back. She can change her size at will, and is a capable fighter when she needs to be, but has never killed or eaten any humans. Instead of gaining nourishment from humans, it seems that she recovers her energy by sleeping for long periods of times."}, {4, "Zenitsu Agatsuma", "https://s4.anilist.co/file/anilistcdn/character/large/b129131-oSy6b121Hetw.png", "Main", "Zenitsu is a coward and is constantly claiming he doesnt have long to live due to the dangerous job of being a Demon Hunter. He is also a bit of a womanizer and likes to hit on girls who he thinks are cute and ask them to marry him -- much to their annoyance. Zenitsu has low self-esteem but still wanted to live up to the expectations others set up for him, but he was constantly afraid and would always cry and run away. He claims he wants to live a modest life where he can be useful to someone."}, {4, "Inosuke Hashibira", "https://s4.anilist.co/file/anilistcdn/character/large/n129130-SJC0Kn1DU39E.jpg", "Main", "Inosuke is extremely short-tempered and proud, and makes a big deal out of fighting opponents stronger than him claiming thats his only hobby. He is a sore loser and is always trying to compete with Tanjiro and provoke him into fighting him, and usually failing. Due to growing up in the mountains by himself Inosuke has trouble interacting with others and only cares about himself most of the time. Though hes often seen with 2 swords in hand, Inosuke is left-handed."}, {5, "Kaname Sudou", "https://s4.anilist.co/file/anilistcdn/character/large/b87219-ufjsOWN7jbpt.jpg", "Main", "He was sent a help request by Kyouda Hiroyuki and joined the Darwins Game. He is only interested in surviving. His sigil allows him to forge any weapon. However, it cant be too complicated."}, {5, "Shuka Karino", "https://s4.anilist.co/file/anilistcdn/character/large/b138795-N2FqVfTrudhO.jpg", "Main", "She is skilled at Darwins game and has 49 wins with no losses. Her sigil is a powerful telekinetic type that allows her to freely control any string-like object."}, {5, "Rein Kashiwagi", "https://s4.anilist.co/file/anilistcdn/character/large/b120241-sH11j8Pj0ERS.jpg", "Main", "A player in Darwins Game who is an information broker. Her sigil is analogous to Laplaces demon, a hypothetical being that knows the exact location and momentum of every atom in the universe."}, {6, "Kazuto Kirigaya", "https://s4.anilist.co/file/anilistcdn/character/large/b36765-IiYQh788PtKs.png", "Main", "Kirito is the main protagonist of the series. He is a solo player, a player who hasnt joined a guild and usually works alone. He is also one of the very few people to have had the privilege to play in the beta testing period of Sword Art Online. His game alias, Kirito, is created by taking the syllables of the first and last Kanji of his real last and first names respectively: (Kirigaya Kazuto)."}, {6, "Eugeo", "https://s4.anilist.co/file/anilistcdn/character/large/b70899-8rr425JDkXc1.jpg", "Main", "Eugeo is the deuteragonist of the first half of the Alicization Arc. Eugeo was a child born in a remote village in the Human Empire of Underworld, where he was assigned the Sacred Task of felling a giant tree with his childhood friend Kirito. After a mishap during an adventure to the mountains resulted in his other childhood friend Alice Zuberg being apprehended by an Integrity Knight for breaking the Taboo Index, while his memories of Kirito were sealed by the administrators of the world, Eugeo spent the following six years focusing on his work to avoid thinking about his lost friend, until Kiritos return to Underworld led to Eugeo building the resolve to embark on a journey to become an Integrity Knight himself to rescue Alice, as well as acquiring the skills to wield the Blue Rose Sword he had retrieved from the mountains where Alice had broken the Taboo Index."}, {6, "Alice Zuberg", "https://s4.anilist.co/file/anilistcdn/character/large/b75450-ylnpHvDUkyfW.jpg", "Main", "The heroine of Sword Art Onlines fourth story arc, Alicization. She is a childhood friend of Kirito and Eugeo. 11 years ago, she broke a taboo in the Taboo Index by entering the Dark Territory by accident, and was later captured by an Integrity Knight of the Axiom Church, the entity that rules the Human World."}, {6, "Asuna Yuuki", "https://s4.anilist.co/file/anilistcdn/character/large/b36828-qoUyIe6c6chj.png", "Main", "Asuna is a friend of Kirito and is a sub-leader of the guild Knights of the Blood (KoB), a medium-sized guild of about thirty players, also called the strongest guild in Aincrad. Being one of the few girls that are in SAO, and even more so that shes extremely pretty, she receives many invitations and proposals. She is a skilled player earning the title Lightning Flash for her extraordinary sword skill that is lightning fast. Her game alias is the same as her real world name."}, {7, "Meliodas", "https://s4.anilist.co/file/anilistcdn/character/large/72921-lUFCsimSUOxQ.png", "Main", "Meliodas the leader of the Seven Deadly Sins and the Sin of Wrath with the symbol of the Dragon. He is the main protagonist of the Nanatsu no Taizai and the owner of the Boar Hat bar."}, {7, "Elizabeth Liones", "https://s4.anilist.co/file/anilistcdn/character/large/b72923-a6q29qhfSAM5.png", "Main", "Elizabeth is the princess of Britannia. She is on a mission to find the seven deadly sins, which are needed to help reclaim the Britannia kingdom."}, {7, "King", "https://s4.anilist.co/file/anilistcdn/character/large/83801-X9xaEsHiVpQ6.jpg", "Main", "King is a member of the Seven Deadly Sins and is the Sin of Sloth with the symbol of the Grizzly. He is also Oslos master. Kings real identity is Fairy King Harlequin, the king of the fairy race."}, {7, "Ban", "https://s4.anilist.co/file/anilistcdn/character/large/b77605-f6lFsdyJS8gP.png", "Main", "Ban is a member of the Seven Deadly Sins and is The Sin of Greed with the symbol of the Fox. He was imprisoned in the Baste Dungeon."}, {7, "Merlin", "https://s4.anilist.co/file/anilistcdn/character/large/89180-KEVDXa6i3DBC.png", "Main", "She is a member of the Seven Deadly Sins. She has the mark of the boars sin of gluttony. She was regarded as the greatest mage in Britannia and is the master of Vivian as well as a teacher of King Arthur of Camelot. She came to the aid of the sins in their battle against Hendirkson but had to leave to tend to the King. She has not officially rejoined the sins but said she would as soon as she was finished with her other business."}, {8, "Licht Bach", "https://s4.anilist.co/file/anilistcdn/character/large/b141532-2AXEveTbFfPU.jpg", "Main", "He was presented at Nanas restaurant as a degenerate lower number, beyond that, fortunately for Hina, he is one of the Legendary Red Barons possessing a powerful force for justicing his path."}, {8, "Hina", "https://s4.anilist.co/file/anilistcdn/character/large/b141533-w7BKIR6ZxHST.jpg", "Main", "She is a solitary girl who lived before with her mom on the mountains. Now she had been traveled around this world looking for The Legendary Red Baron."}, {8, "Jail Murdoch", "https://s4.anilist.co/file/anilistcdn/character/large/b144530-TqjDEwyuYIYe.jpg", "Main", "He is a part of the Althea Royal Guard army, having the rank of squad leader with the title of Lieutenant."}, {9, "Eijun Sawamura", "https://s4.anilist.co/file/anilistcdn/character/large/b22998-U6jtu8OVNjhc.png", "Main", "Sawamura Eijun comes from Nagano where he played baseball in Akagi middle school. He is now a first year high school student of Seidou High School and the roommate of Kuramochi Youichi and Masuko Tooru. He is a left-handed pitcher of the Seidou High School Baseball Team. He was scouted by Takashima Rei in his last year of Junior High."}, {9, "Satoru Furuya", "https://s4.anilist.co/file/anilistcdn/character/large/30266-Tu0UPKz5l5QS.jpg", "Main", "One of the main characters in the manga Ace of Diamond. Furuya is a right-handed power pitcher who throws a fastball and forkball. His main objective is to improve his stamina and control. Eijun Sawamura is his rival."}, {9, "azuya Miyuki", "https://s4.anilist.co/file/anilistcdn/character/large/30267-YsiztSsYpfRq.jpg", "Main", "Miyuki Kazuya is a second year student of Seidou High School and the best catcher of the baseball team. He was already a noted catcher on his first year and nationally known for his skills. Miyuki was scouted by Rei when he was still a first year on his Junior High. As of the current Act, he is a third year."}, {9, "Tesshin Kataoka", "https://s4.anilist.co/file/anilistcdn/character/large/89377-DBWz8WMsYpWk.jpg", "Supporting", "Kataoka Tesshin is known as a strict coach, but he also praises his players and says reassuring things to them from time to time. Everything he does is for the benefit of the team, and he spares no effort in doing so. If he wants his players to make a significant step forward, he himself takes part in training."}}

	clear := `TRUNCATE TABLE anime
	RESTART IDENTITY`
	_, err = db.Exec(clear)
	if err != nil {
		log.Fatalln(err)
	}
	for _, anime := range animes {
		sqlStaement := fmt.Sprintf(`
		INSERT INTO anime (title, description, date, season, imageurl)
		VALUES ('%s','%s','%s','%s','%s')`, anime.title, anime.description, anime.date, anime.season, anime.imageurl)
		_, err = db.Exec(sqlStaement)
		if err != nil {
			log.Fatalln(err)
		}
	}
	clear = `TRUNCATE TABLE genre
	RESTART IDENTITY`
	_, err = db.Exec(clear)
	if err != nil {
		log.Fatalln(err)
	}
	for _, genre := range genres {
		sqlStaement := fmt.Sprintf(`
		INSERT INTO genre (genre) VALUES ('%s')`, genre)
		_, err = db.Exec(sqlStaement)
		if err != nil {
			log.Fatalln(err)
		}
	}
	clear = `TRUNCATE TABLE animegenre
							RESTART IDENTITY`
	_, err = db.Exec(clear)
	if err != nil {
		log.Fatalln(err)
	}
	for _, animegenre := range animegenres {
		sqlStaement := fmt.Sprintf(`
		INSERT INTO animegenre (animeid,genreid) VALUES ('%d', '%d')`, animegenre.animeid, animegenre.genreid)
		_, err = db.Exec(sqlStaement)
		if err != nil {
			log.Fatalln(err)
		}
	}
	clear = `TRUNCATE TABLE character
	RESTART IDENTITY`
	_, err = db.Exec(clear)
	if err != nil {
		log.Fatalln(err)
	}
	for _, character := range characters {
		sqlStaement := fmt.Sprintf(`
		INSERT INTO character (animeid,name,imageurl,role,description) VALUES ('%d', '%s','%s','%s','%s')`, character.animeid, character.name, character.imageurl, character.role, character.description)
		_, err = db.Exec(sqlStaement)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
