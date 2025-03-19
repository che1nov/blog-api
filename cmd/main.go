package main

import (
	"blog-api/internal/db"
	"blog-api/internal/handlers"
	"blog-api/internal/models"
	"blog-api/internal/repositories"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := db.InitDB()

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		slog.Error("Failed to perform auto-migration", "error", err)
		return
	}
	slog.Info("Database migration completed successfully")

	repo := repositories.NewPostRepository(db)
	handler := handlers.NewPostHandler(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/posts", func(r chi.Router) {
		r.Post("/", handler.Create)
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.Get)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})

	slog.Info("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
