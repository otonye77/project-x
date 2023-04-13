package main

import (
	"fmt"
	"net/http"
	"text/template"
)


func Home(w http.ResponseWriter, r *http.Request){
	renderTemplates(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request){
	renderTemplates(w, "about.html")
}

func renderTemplates(w http.ResponseWriter, html string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return 
	}
}