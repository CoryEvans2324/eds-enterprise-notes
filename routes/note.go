package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

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

	jwtUser := middleware.GetUser(r)

	if r.Method == http.MethodGet {
		tmpl.Execute(w, struct{ User *models.JWTUser }{User: jwtUser})
		return
	}

	notetitle := r.FormValue("notetitle")
	notebody := r.FormValue("notecontent")
	noteDateStr := r.FormValue("date")
	noteTimeStr := r.FormValue("time")
	assignedUser := r.FormValue("assigned")
	sharedUserListStr := r.FormValue("sharedUsers")

	var dueDate *time.Time
	if noteDateStr != "" {
		t, err := time.Parse("2006-01-02", noteDateStr)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		dueDate = &t
	}

	if noteTimeStr != "" {
		nTime, err := time.Parse("15:04", noteTimeStr)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// if duedate is nil then use today
		if dueDate == nil {
			t := time.Now()
			td := time.Date(t.Year(), t.Month(), t.Day(), nTime.Hour(), nTime.Minute(), 0, 0, time.Local)
			dueDate = &td
		} else {
			// add the two together
			t := time.Date(dueDate.Year(), dueDate.Month(), dueDate.Day(), nTime.Hour(), nTime.Minute(), 0, 0, time.Local)
			dueDate = &t
		}
	}

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

	user, _ := database.Mgr.GetUserByID(jwtUser.UserID)
	note := &models.Note{
		Name:          notetitle,
		Content:       notebody,
		Status:        "In progress",
		Owner:         user,
		DueDate:       dueDate,
		DelegatedUser: delegatedUser,
		SharedUsers:   permissions,
	}

	note, err = database.Mgr.CreateNote(note)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/note/%d", note.ID), http.StatusFound)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	note, err := database.Mgr.GetNoteByID(uint(id))
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
		User *models.JWTUser
		Note *models.Note
	}{
		User: middleware.GetUser(r),
		Note: note,
	})
}
