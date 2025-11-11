package main

import (
	"fmt"
	"net/http"

	"github.com/bielbergamini/go_web_api/internal/handler"
)

func main() {
	http.HandleFunc("/status", handler.StatusHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
