package templates

import (
	"fmt"
	"html/template"
)

type Element struct {
	filename string
	parent   *Element
}

type TemplateList struct {
	Length int
	Head   *Element
}

// Creates a new TemplateList
func NewTemplateList(filenames ...string) *TemplateList {
	l := &TemplateList{}
	for _, fn := range filenames {
		l.Append(fn)
	}
	return l
}

// Appends a filename of a template to the front of the linked list
func (tl *TemplateList) Append(filename string) {
	ele := &Element{
		filename: filename,
	}

	if tl.Head == nil {
		tl.Head = ele
	} else {
		ele.parent = tl.Head
		tl.Head = ele
	}

	tl.Length++
}

// Clones the list and appends a filename
func (tl *TemplateList) Extend(filename string) *TemplateList {
	newList := *tl
	newList.Append(filename)
	return &newList
}

func (tl *TemplateList) CreateHtmlTemplate() (*template.Template, error) {
	if tl.Head == nil {
		return nil, fmt.Errorf("template list has no elements")
	}

	files := make([]string, tl.Length)

	cur := tl.Head
	for i := 0; i < tl.Length; i++ {
		// TODO: load templates folder location from config
		files[i] = "web/templates/" + cur.filename
		cur = cur.parent
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, fmt.Errorf("error parsing files")
	}

	return tmpl, nil
}
