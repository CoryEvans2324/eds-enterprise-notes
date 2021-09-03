package main

import (
	"io/ioutil"
	"log"

	"github.com/CoryEvans2324/eds-enterprise-notes/config"
	"github.com/CoryEvans2324/eds-enterprise-notes/database"
)

func init() {
	cfgData, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	config.LoadConfig(cfgData)
}

func checkErrorNil(err error) {
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
}

func main() {
	cfg := config.Get()

	err := database.CreateDatabaseManager(cfg.Database.DataSourceName())
	checkErrorNil(err)

	createFakeUsers()
	createFakeNotes()
}

func createFakeUsers() {
	users := []string{
		"cory",
		"Alleen",
		"Demeter",
		"Bamby",
		"Brenda",
		"Ophelie",
		"Tobe",
		"Nada",
		"Fey",
		"Janeczka",
		"Merissa",
		"Nancy",
		"Gavra",
		"Jessika",
		"Charisse",
		"Wynn",
		"Linnet",
		"Verna",
		"Cacilie",
		"Moina",
		"Ardyth",
	}

	for i := 0; i < len(users); i += 2 {
		database.Mgr.CreateUser(users[i], "password123")
	}
}

func createFakeNotes() {

}
