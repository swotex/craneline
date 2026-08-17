package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"octopus/internal/http/handlers"
)

func NewRouter(imageHandler *handlers.ImageHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/images", func(r chi.Router) {
			r.Get("/all", imageHandler.ListImages)
			r.Get("/{id}", imageHandler.GetImage)
			r.Post("/new", imageHandler.CreateFullImage)
			r.Route("/versions/{id}", func(r chi.Router) {
				r.Post("/new", imageHandler.AddVersion)
				r.Get("/all", imageHandler.ListVersions)
			})
		})
	})

	return r
}
