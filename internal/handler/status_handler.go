package handler

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := StatusResponse{
		Status:  "ok",
		Message: "API is running successfully!",
	}
	json.NewEncoder(w).Encode(response)
}
