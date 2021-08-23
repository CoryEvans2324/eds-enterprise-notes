package routes

import (
	"net/http"

	"github.com/CoryEvans2324/eds-enterprise-notes/database"
)

func DebugResetDB(w http.ResponseWriter, r *http.Request) {
	database.Mgr.DropUserTable()
	database.Mgr.CreateUserTable()
	database.Mgr.DropPermissionTable()
	database.Mgr.CreatePermissionTable()
	database.Mgr.DropNoteTable()
	database.Mgr.CreateNoteTable()

	w.Write([]byte("OK\n"))
}

func DebugCreateDummyUsers(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte("OK\n"))
}
