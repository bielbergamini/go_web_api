package http

import (
	"database/sql"
	"net/http"

	"go_web_api/internal/infrastructure/db"
	"go_web_api/internal/infrastructure/http/handler"
)

func NewRouter(database *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	repo := db.NewUserRepository(database)
	userHandler := handler.NewUserHandler(repo)

	mux.HandleFunc("/users", userHandler.CreateUser)
	mux.HandleFunc("/users/", userHandler.GetUser)

	return mux
}
