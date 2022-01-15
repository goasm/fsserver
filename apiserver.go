package main

import (
	"fmt"
	"net/http"
)

type APIHandler struct {
}

func APIServer() APIHandler {
	return APIHandler{}
}

func (h APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL)
}
