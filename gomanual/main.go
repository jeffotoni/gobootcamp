package main

import (
	"net/http"
)

var (
	HTTP_PORT = "0.0.0.0:8181"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir("./manual"))))
	println("Run Server: http://localhost:8181")
	http.ListenAndServe(HTTP_PORT, mux)
}
