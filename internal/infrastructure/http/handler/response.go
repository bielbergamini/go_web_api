package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}


func writeError(w http.ResponseWriter, r *http.Request, status int, message string) {
	hlog.FromRequest(r).Error().Int("status", status).Msg(message)
	writeJSON(w, status, map[string]string{"error": message})
}
