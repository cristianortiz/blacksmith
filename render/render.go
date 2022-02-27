package render

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set //support for jet
	Session    *scs.SessionManager
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

//defaultData set the init values if for a template view to be rendered
//data is extracted for the Render type receiver
func (rd *Render) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = rd.Secure
	td.ServerName = rd.ServerName
	td.Port = rd.Port
	//check if user session variable  exists to authenticate a user
	if rd.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	return td
}

//Render show a page switching between supported templates engines,view is the name of the file to be
//rendered, variables are config option for JetViews and data are the dynamic content to render.
// both are interfaces because can be anything
func (rd *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	//use different types of renderers
	switch strings.ToLower(rd.Renderer) {
	case "go":
		return rd.GoPage(w, r, view, data)

	case "jet":
		return rd.JetPage(w, r, view, variables, data)
	}
	return errors.New("no rendering engine specified")

}

//GoPage render a view using Go Templates
func (rd *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
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

//JetPage renders a template using the Jet templating engine
func (rd *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	//at first variables and data params must be process, because variables are the Jet parameters and data
	//are the templateData content, JetPage() need both  in the correct format to render the view

	//empty jet.VarMap variable
	var vars jet.VarMap
	//check if variables param is not empty
	if variables == nil {
		//init the above empty jet.VarMap
		vars = make(jet.VarMap)
	} else {
		//if variables have something, cast it to jet.VarMap Format to work with jet
		vars = variables.(jet.VarMap)
	}
	// data param (templateData content) must be procesed to the correct fomat
	td := &TemplateData{} //empty TemplateData struct
	//if data param (the templateData content) is empty
	if data != nil {
		//cast data to TemplateData type, the Format to wor with it
		td = data.(*TemplateData)
	}

	td = rd.defaultData(td, r)
	//get the template file by their name
	template, err := rd.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}
	//render the template
	err = template.Execute(w, vars, td)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
