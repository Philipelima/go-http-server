package main

import (
	"fmt"
	"net/http"

	"github.com/philipelima/http-server/internal/server"
)

func hello(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "<h1>Hello, Server</h1>")
}

func main() {

	server := server.Server{}

	server.SetDirectory("./static")
	server.SetSettings("127.0.0.1", "4000")

	server.Get("/hello-server", hello)

	server.Run()
}
