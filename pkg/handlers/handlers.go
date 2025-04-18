package handlers

import (
	"log"
	"net/http"

	"github.com/Priyansh-Kotak/booking-system/pkg/config"
	"github.com/Priyansh-Kotak/booking-system/pkg/models"
	"github.com/Priyansh-Kotak/booking-system/pkg/render"
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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	log.Println("local ip address", remoteIP)

	stringMap := make(map[string]string)
	stringMap["testing"] = "Hello , developers"

	render.RenderTemplets(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page render
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)

	StringMap["test"] = "hey -- Priyansh this side"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP
	log.Println("Printing about ip ", remoteIP)

	render.RenderTemplets(w, "about.page.html", &models.TemplateData{
		StringMap: StringMap,
	})
}
