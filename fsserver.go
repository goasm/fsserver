package fsserver

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
)

func FSServer(root string) http.Handler {
	return &fsServer{
		root: root,
	}
}

type Result struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type fsServer struct {
	root string
}

func (s *fsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.Get(w, r)
	case http.MethodPost:
		s.Post(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *fsServer) Get(w http.ResponseWriter, r *http.Request) {
	rel := r.URL.Path
	abs := path.Join(s.root, rel)
	http.ServeFile(w, r, abs)
}

func (s *fsServer) Post(w http.ResponseWriter, r *http.Request) {
	rel := r.URL.Path
	abs := path.Join(s.root, rel)
	result := Result{}
	defer json.NewEncoder(w).Encode(result)
	err := SaveFile(r.Body, abs)
	if err != nil {
		log.Printf("{error: SaveFile} %s", err.Error())
		result.Success = false
		result.Error = "failed to save file"
		return
	}
	result.Success = true
}
