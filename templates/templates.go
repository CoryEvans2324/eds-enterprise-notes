package templates

import (
	"html/template"
	"net/http"
)

var templates = make(map[string]*template.Template)

// Creates and saves a template by name
func CreateTemplate(name string, tmplList TemplateList) {
	templates[name] = tmplList.CreateHtmlTemplate()
}

// Gets a template by name
func GetTemplate(name string) *template.Template {
	return templates[name]
}

// Executes a template on a http.ResponseWriter
func RenderTemplate(w http.ResponseWriter, name string, pageData *interface{}) {
	tmpl := GetTemplate(name)

	tmpl.Execute(w, pageData)
}
