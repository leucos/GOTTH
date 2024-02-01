package handlers

import (
	"goth/internal/templates"
	"net/http"
)

type HomeHander struct{}

func NewHomeHandler() *HomeHander {
	return &HomeHander{}
}

func (h *HomeHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Index()
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}