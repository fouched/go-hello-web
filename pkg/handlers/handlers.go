package handlers

import (
	"github.com/fouched/go-course/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplage(w, "home.page.gohtml")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplage(w, "about.page.gohtml")
}
