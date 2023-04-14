package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"main.com/mymodule/pkg/config"
)

var functions = template.FuncMap {

}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func RenderTemplates(w http.ResponseWriter, html string) {
	tc := app.TemplateCache

	t, ok := tc[html]
	if !ok {
		log.Fatal("Could not get template from template cacje")
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error){
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = t
	}
	return myCache, nil
}
