package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

//to render a page switching between supported templates engines, variables and data are interfaces because can be anything
func (rd *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	//use different types of renderers
	switch strings.ToLower(rd.Renderer) {
	case "go":
		return rd.GoPage(w, r, view, variables, data)

	case "jet":
	}
	return nil

}

//GoPage
func (rd *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	//parse  html template file
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", rd.RootPath, view))
	if err != nil {
		return err
	}
	//reference to templateData struct
	td := &TemplateData{}
	//if data to show in the html page is not null
	if data != nil {
		td = data.(*TemplateData)
	}
	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}
	return nil
}
