package main

import (
	"encoding/json"
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
		defer json.NewEncoder(w).Encode(result)
		cwd := r.FormValue("cwd")
		name := r.FormValue("name")
		file, _, err := r.FormFile("file")
		if err != nil {
			result = Result{Success: false, Error: "failed to parse form"}
			return
		}
		// join destination file path
		root, err := os.Getwd()
		if err != nil {
			result = Result{Success: false, Error: "failed to prepare file"}
			return
		}
		destPath := filepath.Join(root, cwd, name)
		err = SaveFile(file, destPath)
		if err != nil {
			result = Result{Success: false, Error: "failed to save file"}
			return
		}
		result = Result{Success: true}
	})
	return mux
}
