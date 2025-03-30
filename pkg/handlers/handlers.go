package handlers

import (
	"net/http"

	"github.com/Priyansh-Kotak/udemy-course-project/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplets(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplets(w, "about.page.html")
}
