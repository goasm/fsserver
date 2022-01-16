package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Result struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func APIServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		result := Result{}
		if err != nil {
			result = Result{Success: false, Error: err.Error()}
			return
		}
		dest, err := os.OpenFile("dest", os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			result = Result{Success: false, Error: err.Error()}
			return
		}
		_, err = io.Copy(dest, file)
		if err != nil {
			result = Result{Success: false, Error: err.Error()}
			return
		}
		result = Result{Success: true}
		json.NewEncoder(w).Encode(result)
	})
	return mux
}
