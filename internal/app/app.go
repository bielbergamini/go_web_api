package app

import (
	"fmt"
	"net/http"

	apihttp "go_web_api/internal/infrastructure/http"
)

func Run() error {
	router := apihttp.NewRouter()

	fmt.Println("Servidor running on http://localhost:8080")
	return http.ListenAndServe(":8080", router)
}
