package http

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}