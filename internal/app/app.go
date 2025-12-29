package app

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"go_web_api/internal/config"
	"go_web_api/internal/infrastructure/db"
	apihttp "go_web_api/internal/infrastructure/http"
)

func Run() error {
	cfg := config.Load()

	database, err := db.NewPostgresConnection(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer database.Close()

	router := apihttp.NewRouter(database)

	addr := ":" + cfg.ServerPort
	log.Info().Msgf("Server running at http://localhost%s", addr)

	return http.ListenAndServe(addr, router)
}
