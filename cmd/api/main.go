package main

import (
	"go_web_api/internal/app"
	"go_web_api/internal/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	logger.Initialize()
	if err := app.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run app")
	}
}
