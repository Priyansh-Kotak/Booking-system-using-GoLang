package handlers

import (
	"net/http"

	"github.com/Priyansh-Kotak/udemy-course-project/pkg/config"
	"github.com/Priyansh-Kotak/udemy-course-project/pkg/models"
	"github.com/Priyansh-Kotak/udemy-course-project/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creats a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handler
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page render
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["testing"] = "Hello , developers"

	render.RenderTemplets(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page render
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplets(w, "about.page.html", &models.TemplateData{})
}
