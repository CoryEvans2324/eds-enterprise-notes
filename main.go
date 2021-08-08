package main

import (
	"github.com/CoryEvans2324/eds-enterprise-notes/templates"
)

func main() {
	templates.CreateTemplate("index", "index.html", "base.layout.html")
}
