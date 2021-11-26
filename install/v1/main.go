package main

import (
	"embed"
	"net/http"
)

var (
	HTTP_PORT = "0.0.0.0:8181"
)

//go:embed manual

var res embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/",
		http.FileServer(http.FS(res))))
	println("Run Server: http://localhost:8181/manual")
	http.ListenAndServe(HTTP_PORT, mux)
}
