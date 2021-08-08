package routes

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/index.html", "web/base.layout.html")

	if err != nil {
		log.Fatalf("Index: %v\n", err)
	}

	tmpl.Execute(w, nil)
}
