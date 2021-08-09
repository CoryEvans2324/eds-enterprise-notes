package routes

import (
	"html/template"
	"log"
	"net/http"
)

func UserSignIn(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/user/signin.html", "web/base.layout.html")

	if err != nil {
		log.Fatalf("UserSignIn: %v\n", err)
	}

	tmpl.Execute(w, nil)
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/user/create.html", "web/base.layout.html")

	if err != nil {
		log.Fatalf("UserSignUp: %v\n", err)
	}

	tmpl.Execute(w, nil)
}
