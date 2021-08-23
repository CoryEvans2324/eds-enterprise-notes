package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/CoryEvans2324/eds-enterprise-notes/database"
	"github.com/CoryEvans2324/eds-enterprise-notes/middleware"
	"github.com/CoryEvans2324/eds-enterprise-notes/models"
	"github.com/gorilla/mux"
)

type sharedUser struct {
	Username string `json:"username"`
	Editor   bool   `json:"editor"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/note/create.html", "web/base.layout.html")

	user := middleware.GetUser(r)

	if r.Method == http.MethodGet {
		tmpl.Execute(w, struct{ User *models.User }{User: user})
		return
	}

	notetitle := r.FormValue("notetitle")
	notebody := r.FormValue("notecontent")
	// noteDateStr := r.FormValue("date")
	// noteTimeStr := r.FormValue("time")
	assignedUser := r.FormValue("assigned")
	sharedUserListStr := r.FormValue("sharedUsers")

	var sharedUsers = make([]sharedUser, 0)

	err := json.Unmarshal([]byte(sharedUserListStr), &sharedUsers)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var delegatedUser *models.User
	if assignedUser != "" {
		delegatedUser, err = database.Mgr.GetUserByUsername(assignedUser)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	}

	var permissions = make([]models.Permission, 0)
	for _, su := range sharedUsers {
		sum, err := database.Mgr.GetUserByUsername(su.Username)
		if err != nil {
			log.Println(err)
			continue
		}

		var perm string
		if su.Editor {
			perm = "editor"
		} else {
			perm = "viewer"
		}
		permissions = append(permissions, models.Permission{
			Permission: perm,
			User:       *sum,
		})
	}

	note := models.Note{
		Name:          notetitle,
		Content:       notebody,
		Status:        "In progress",
		Owner:         user,
		DelegatedUser: delegatedUser,
		SharedUsers:   permissions,
	}

	noteID, err := database.Mgr.CreateNote(note)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/%d", noteID), http.StatusFound)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	note, err := database.Mgr.GetNoteByID(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tmpl, err := template.ParseFiles("web/note/note.html", "web/base.layout.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, struct {
		User *models.User
		Note *models.Note
	}{
		User: middleware.GetUser(r),
		Note: note,
	})
}
