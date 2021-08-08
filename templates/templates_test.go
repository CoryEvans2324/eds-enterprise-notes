package templates

import (
	"bytes"
	"html/template"
	"net/http"
	"testing"
)

// ResponseRecorder is an implementation of http.ResponseWriter that
// records its mutations for later inspection in tests.
type ResponseRecorder struct {
	Code      int           // the HTTP response code from WriteHeader
	HeaderMap http.Header   // the HTTP response headers
	Body      *bytes.Buffer // if non-nil, the bytes.Buffer to append written data to
	Flushed   bool
}

// NewRecorder returns an initialized ResponseRecorder.
func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		HeaderMap: make(http.Header),
		Body:      new(bytes.Buffer),
	}
}

// Header returns the response headers.
func (rw *ResponseRecorder) Header() http.Header {
	return rw.HeaderMap
}

// Write always succeeds and writes to rw.Body, if not nil.
func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
	if rw.Body != nil {
		rw.Body.Write(buf)
	}
	if rw.Code == 0 {
		rw.Code = http.StatusOK
	}
	return len(buf), nil
}

// WriteHeader sets rw.Code.
func (rw *ResponseRecorder) WriteHeader(code int) {
	rw.Code = code
}

// Flush sets rw.Flushed to true.
func (rw *ResponseRecorder) Flush() {
	rw.Flushed = true
}

func TestTemplates(t *testing.T) {
	base := NewTemplateList("base.layout.html")
	index := base.Extend("index.html")

	if index.Head.filename != "index.html" {
		t.Fatalf("Extend failed")
	}

	var err error

	CreateTemplate("test", index)

	empty := NewTemplateList()
	_, err = empty.CreateHtmlTemplate()
	if err == nil {
		t.Errorf("Empty template list didn't return an error")
	}

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>
`
	tmpl, err := template.New("test").Parse(tpl)
	if err != nil {
		t.Errorf(err.Error())
	}

	templates["test"] = tmpl

	rw := NewRecorder()
	err = RenderTemplate(rw, "test", noItems)
	if err != nil {
		t.Errorf(err.Error())
	}
	RenderTemplate(rw, "doesnt exist", noItems)
}
