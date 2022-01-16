package main

import (
	"fmt"
	"net/http"
)

func APIServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("2 ", r.URL)
	})
	return mux
}
