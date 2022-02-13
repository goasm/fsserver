package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	fss "github.com/goasm/fsserver"
)

var (
	host string
	port string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: fsserver [OPTION...] PATH")
		flag.PrintDefaults()
	}
	flag.StringVar(&host, "a", "0.0.0.0", "address to use")
	flag.StringVar(&port, "p", "8080", "port to bind to")
}

func printServerInfo(root string) {
	fmt.Println("Serving path:", root)
	ip, err := fss.LocalIP()
	if host == "0.0.0.0" && err == nil {
		fmt.Println("Available on:", "http://"+net.JoinHostPort(ip, port))
	}
}

func main() {
	flag.Parse()
	root, _ := filepath.Abs(flag.Arg(0))
	addr := net.JoinHostPort(host, port)
	server := fss.Compose(fss.FSServer(root), []fss.Middleware{fss.LogRequest, fss.JsonPostResponse})
	http.Handle("/", server)
	printServerInfo(root)
	log.Fatal(http.ListenAndServe(addr, nil))
}
