package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Success bool `json:"success"`
}

func APIServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Result{Success: true})
	})
	return mux
}
