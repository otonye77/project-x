package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplatesTest(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplates(w http.ResponseWriter, t string) {
	var html *template.Template
	var err error
	_, inMap := tc[t]
	if !inMap {
		err = createTemplateCache(t)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
	} else {
		fmt.Println("Reading cached template")
	}
	html = tc[t]
	err = html.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.html",
	}
	html, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = html
	return nil
}
