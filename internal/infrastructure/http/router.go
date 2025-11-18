package http

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// basic route for testing
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API running"))
	})

	return mux
}
