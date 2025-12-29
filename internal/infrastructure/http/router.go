package http

import (
	"database/sql"
	"go_web_api/internal/config"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"go_web_api/internal/auth"
	"go_web_api/internal/infrastructure/db"
	"go_web_api/internal/infrastructure/http/handler"
)

func NewRouter(database *sql.DB, cfg config.Config) http.Handler {
	mux := chi.NewRouter()

	mux.Use(hlog.NewHandler(log.Logger))
	mux.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	mux.Use(middleware.Recoverer)

	authenticator := auth.New(cfg.JWTSecret)

	repo := db.NewUserRepository(database)
	userHandler := handler.NewUserHandler(repo, authenticator)

	// Public routes
	mux.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})
	mux.Post("/login", userHandler.Login)
	mux.Post("/users", userHandler.CreateUser)

	// Protected routes
	mux.Group(func(r chi.Router) {
		r.Use(authenticator.Middleware)
		r.Get("/users/{id}", userHandler.GetUser)
		r.Put("/users/{id}", userHandler.UpdateUser)
		r.Delete("/users/{id}", userHandler.DeleteUser)
	})

	return mux
}
