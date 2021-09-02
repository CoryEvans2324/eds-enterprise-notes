package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/CoryEvans2324/eds-enterprise-notes/database"
	"github.com/CoryEvans2324/eds-enterprise-notes/middleware"
	"github.com/CoryEvans2324/eds-enterprise-notes/models"
	"github.com/gorilla/mux"
)

func UserView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := database.Mgr.GetUserByUsername(username)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	tmpl, _ := template.ParseFiles("web/user/user.html", "web/base.layout.html")
	tmpl.Execute(w, struct{ User *models.User }{User: user})
}

func UserSignIn(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/user/signin.html", "web/base.layout.html")

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := database.Mgr.GetUserByUsername(username)
	if err != nil {
		log.Println(err)
		tmpl.Execute(w, nil)
		return
	}

	if !database.CheckPasswordWithHash(password, user.PasswordHash) {
		log.Println("password incorrect")
		tmpl.Execute(w, nil)
		return
	}
	middleware.SetUser(w, &models.JWTUser{
		UserID: user.ID,
	})

	http.Redirect(w, r, "/", http.StatusFound)
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/user/create.html", "web/base.layout.html")

	if err != nil {
		log.Fatalf("UserSignUp: %v\n", err)
	}
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := database.Mgr.CreateUser(username, password)
	if err != nil {
		tmpl.Execute(w, nil)
		return
	}

	middleware.SetUser(w, &models.JWTUser{
		UserID: user.ID,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func UserSignOut(w http.ResponseWriter, r *http.Request) {
	middleware.SetUser(w, nil)
	http.Redirect(w, r, "/", http.StatusFound)
}

func UserSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	usernames, usernamePresent := query["username"]
	searchFor, searchForPresent := query["for"]

	if !usernamePresent || !searchForPresent || len(usernames) == 0 {
		return
	}

	users, err := database.Mgr.SearchForUsername(usernames[0])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	templateData := struct {
		Usernames []struct {
			Username string
		}
	}{}

	for _, v := range users {
		templateData.Usernames = append(templateData.Usernames, struct{ Username string }{v})
	}

	var tmpl *template.Template
	switch searchFor[0] {
	case "assignment":
		tmpl, _ = template.ParseFiles("web/user/assignedUserSearch.html")

	case "sharing":
		tmpl, _ = template.ParseFiles("web/user/sharedUserSearch.html")

	default:
		return
	}

	tmpl.Execute(w, templateData)
}
