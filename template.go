package tmpl

import (
	"html/template"
	"io"
	"log"
	"time"
)

const defaultLayoutValue = "application_layout"

// Template struct
type Template struct {
	tmpl *template.Template

	defaultLayout string
}

// NewTemplate creates a new instance of the template library
func NewTemplate(opts ...Optioner) (*Template, error) {
	defaultLayout := defaultLayoutValue
	for _, opt := range opts {
		if opt.getName() == optionLayoutName {
			defaultLayout = opt.getValue()
		}
	}

	tmpl, err := template.ParseGlob("template/shared/*")
	if err != nil {
		return nil, err
	}

	tmpl, err = tmpl.ParseGlob("template/layout/*")

	return &Template{tmpl: tmpl, defaultLayout: defaultLayout}, nil
}

// AddFiles to the current template
func (t *Template) AddFiles(files ...string) (err error) {
	var prefixedFiles []string
	for _, f := range files {
		prefixedFiles = append(prefixedFiles, "template/"+f)
	}

	t.tmpl, err = t.tmpl.ParseFiles(prefixedFiles...)

	return err
}

// Clone returns a duplicate of the template, including all associated templates.
// The actual representation is not copied, but the name space of associated templates is,
// so further calls to Parse in the copy will add templates to the copy but not to the original.
// Clone can be used to prepare common templates and use them with variant definitions for
// other templates by adding the variants after the clone is made.
func (t *Template) Clone() (*Template, error) {
	tmpl, err := t.tmpl.Clone()
	return &Template{tmpl: tmpl, defaultLayout: t.defaultLayout}, err
}

// Execute the template with the default layout
func (t Template) Execute(w io.Writer, data interface{}) error {
	return t.ExecuteLayout(w, t.defaultLayout, data)
}

// ExecuteLayout executes the template with a given layout name
func (t Template) ExecuteLayout(w io.Writer, layoutName string, data interface{}) error {
	start := time.Now()

	err := t.tmpl.ExecuteTemplate(w, layoutName, data)
	if err != nil {
		log.Println(err)
	}

	elapsed := time.Since(start)
	log.Printf("Completed Layout Render in %s\n", elapsed)

	return err
}
