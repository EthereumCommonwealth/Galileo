package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

//could be both constants or variables. (testing purposes)
var (
	LayoutDir string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View  {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View {
		Template: t,
		Layout: layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View ) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}
type View struct {
	Template *template.Template
	Layout string
}