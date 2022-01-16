package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Result struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func APIServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		result := Result{}
		cwd := r.FormValue("cwd")
		name := r.FormValue("name")
		file, _, err := r.FormFile("file")
		// join destination file path
		root, err := os.Getwd()
		if err != nil {
			result = Result{Success: false, Error: err.Error()}
			return
		}
		destPath := filepath.Join(root, cwd, name)
		if err != nil {
			result = Result{Success: false, Error: err.Error()}
			return
		}
		dest, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0644)
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
