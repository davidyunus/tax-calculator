package httpserver

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func (hs *HTTPServer) compileRouter() chi.Router {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	// Tax API
	r.Route("/v1", func(r chi.Router) {
		r.Post("/tax", hs.taxController.Create)
		r.Get("/taxes", hs.taxController.FindAll)
		r.Get("/tax/{taxId}", hs.taxController.FindByKeys)
		r.Get("/tax", hs.taxController.FindByQuery)
		r.Put("/tax/{taxId}", hs.taxController.Update)
	})

	return r
}
