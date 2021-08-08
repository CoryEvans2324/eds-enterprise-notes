package templates

import (
	"html/template"
)

func CreateTemplate(name string, filenames ...string) {
	files := make([]string, len(filenames))

	for i := 0; i < len(filenames); i++ {
		files[i] = "web/templates/" + filenames[i]
	}

	template.New(name).ParseFiles(files...)
}
