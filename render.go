package main

import (
	"fmt"
	"net/http"
	"text/template"
)


func renderTemplates(w http.ResponseWriter, html string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates", err)
		return 
	}
}