package app

import (
	"fmt"
	"net/http"

	"go_web_api/internal/config"
	apihttp "go_web_api/internal/infrastructure/http"
)

func Run() error {
	cfg := config.Load()

	router := apihttp.NewRouter()

	addr := ":" + cfg.ServerPort

	fmt.Printf("Server running on http://localhost%s\n", addr)

	return http.ListenAndServe(addr, router)
}
