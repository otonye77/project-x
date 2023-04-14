package handlers

import (
	"net/http"

	"main.com/mymodule/pkg/config"
	"main.com/mymodule/pkg/models"
	"main.com/mymodule/pkg/render"
)


var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplates(w, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"
	render.RenderTemplates(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

