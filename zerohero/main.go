package main

import (
	"log"
	"net/http"
	"runtime"
)

type ZeroHero struct {
	ID         string     `json:"id"`
	Response   string     `json:"response"`
	Name       string     `json:"name"`
	Powerstats Powerstats `json:"powerstats"`
	Biography  Biography  `json:"biography"`
	Work       Work       `json:"work"`
	Image      Image      `json:"image"`
}
type Powerstats struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}
type Biography struct {
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceOfBirth    string   `json:"place-of-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}
type Work struct {
	Occupation string `json:"occupation"`
	Base       string `json:"base"`
}
type Image struct {
	URL string `json:"url"`
}

// JWT
// CORS
// LOGGER
// RATE LIMIT
// CACHE
// INSTRUMENTACAO
// BANCO DE DADOS
// GERAÇÃO DE LOGS SAÍDA

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong😍"))
		})

	mux.HandleFunc("/api/v1/zerohero", Service)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Println("\033[1;44mRunning on http://0.0.0.0:8080 (Press CTRL+C to quit)\033[0m")
	log.Fatal(server.ListenAndServe())
}

func Service(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		Post(w, r)
	case http.MethodGet:
		Get(w, r)
	case http.MethodDelete:
		Delete(w, r)
	case http.MethodPut:
		Put(w, r)
	default:
		http.NotFound(w, r)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"Post"}`))
}
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"msg":"Get"}`))
}
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"msg":"Delete"}`))
}
func Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"msg":"Put"}`))
}
