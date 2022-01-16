package main

import (
	"log"
	"net"
	"net/http"
)

// LocalIP returns the IP address of the localhost
func LocalIP() (string, error) {
	addr, err := net.ResolveUDPAddr("udp", "1.2.3.4:1")
	if err != nil {
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	ip, _, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return "", err
	}
	return ip, nil
}

type Middleware func(http.Handler) http.Handler

func Compose(h http.Handler, mws []Middleware) http.Handler {
	for i := 0; i < len(mws); i++ {
		h = mws[i](h)
	}
	return h
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func JsonResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
