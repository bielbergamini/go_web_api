package app

import (
	"fmt"
	"log"
	"net/http"

	"go_web_api/internal/config"
	"go_web_api/internal/infrastructure/db"
	apihttp "go_web_api/internal/infrastructure/http"
)

func Run() error {
	cfg := config.Load()

	database, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	_ = database 

	router := apihttp.NewRouter()

	addr := ":" + cfg.ServerPort
	fmt.Printf("Server running at http://localhost%s\n", addr)

	return http.ListenAndServe(addr, router)
}
