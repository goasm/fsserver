package fsserver

import (
	"net/http"
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

func (s *fsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
