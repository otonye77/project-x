package handlers

import (
	"net/http"

	"main.com/mymodule/pkg/render"
)


func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplates(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplates(w, "about.html")
}

