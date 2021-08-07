package main

import (
	"github.com/CoryEvans2324/eds-enterprise-notes/templates"
)

func main() {
	baseLayout := templates.NewTemplateList("base.layout.html")
	indexTemplate := baseLayout.Extend("index.html")

	templates.CreateTemplate("index", *indexTemplate)
}
