package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct que representa a resposta JSON
type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Handler da rota /status
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := StatusResponse{
		Status:  "ok",
		Message: "API is running successfully!",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/status", statusHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
