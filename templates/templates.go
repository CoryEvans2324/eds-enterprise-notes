package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = make(map[string]*template.Template)

// Creates and saves a template by name
func CreateTemplate(name string, tmplList *TemplateList) error {
	tmpl, err := tmplList.CreateHtmlTemplate()
	if err != nil {
		return fmt.Errorf("CreateTemplate: %v", err)
	}

	templates[name] = tmpl
	return nil
}

// Gets a template by name
func GetTemplate(name string) (*template.Template, error) {
	value, ok := templates[name]
	if !ok {
		return nil, fmt.Errorf("no template with name %s", name)
	}

	return value, nil
}

// Executes a template on a http.ResponseWriter
func RenderTemplate(w http.ResponseWriter, name string, pageData interface{}) error {
	tmpl, err := GetTemplate(name)
	if err != nil {
		return fmt.Errorf("RenderTemplate: %v", err)
	}

	return tmpl.Execute(w, pageData)
}
