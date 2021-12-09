package main

import (
	"log"
	"net/http"
	"runtime"

	ha "interno/handler"
	mw "interno/middleware"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong😍"))
		})

	mux.HandleFunc("/api/v1/zerohero", mw.Use(ha.Service, mw.Logger()))
	mux.HandleFunc("/", mw.Use(ha.Service, mw.Logger()))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Println("\033[1;44mRunning on http://0.0.0.0:8080 (Press CTRL+C to quit)\033[0m")
	log.Fatal(server.ListenAndServe())
}
