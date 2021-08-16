package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/CoryEvans2324/eds-enterprise-notes/middleware"
	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/index.html", "web/base.layout.html")

	if err != nil {
		log.Fatalf("Index: %v\n", err)
	}

	user := middleware.GetUser(r)

	tmpl.Execute(w, struct{ User *models.User }{User: user})
}
