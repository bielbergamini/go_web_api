package http

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_web_api/internal/infrastructure/db"
	"go_web_api/internal/infrastructure/http/handler"
)

func NewRouter(database *sql.DB) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	repo := db.NewUserRepository(database)
	userHandler := handler.NewUserHandler(repo)

	mux.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	return mux
}
