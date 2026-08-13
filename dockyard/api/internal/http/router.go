package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"octopus/internal/http/handlers"
)

func NewRouter(imageHandler *handlers.ImageHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/images", func(r chi.Router) {
			r.Get("/all", imageHandler.ListImages)
			r.Post("/new", imageHandler.CreateFullImage)
			r.Get("/{id}", imageHandler.GetImage)
		})
	})

	return r
}
