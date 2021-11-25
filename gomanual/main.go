package main

import (
	"net/http"
)

var (
	HTTP_PORT = "0.0.0.0:8181"
)

////go:embed manual
//var res embed.FS

func main() {
	mux := http.NewServeMux()
	//fs := http.FileServer(http.FS(res))
	//mux.Handle("/manual/", http.StripPrefix("/", fs))
	mux.Handle("/manual/", http.StripPrefix("/manual",
		http.FileServer(http.Dir("./manual"))))
	println("Run Server ", HTTP_PORT)
	http.ListenAndServe(HTTP_PORT, mux)
}
