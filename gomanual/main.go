package main

import (
	"embed"
	"log"
	"net/http"

	"pkg/util"
)

// Variaveis globais
var (
	version     string
	PORT        = "8080"
	HTTP_PORT   = "0.0.0.0:" + PORT
	HOST_DOCKER = "http://localhost:" + PORT + "/manual"
	HOST        = "http://localhost:" + PORT
)

// Embed é um recurso nativo que temos
// em Go para que possamo colocar dentro do
// binário todo conteúdo que encontra-se
// no diretório manual
//go:embed manual

// variavel criado do ditpo embed
var res embed.FS

func main() {
	// usamos go build -ldflags="-X main.version=docker"
	// aqui pegamos o valor passado na compilacao
	// para escolhermos qual forma iremos servir
	// o manual
	switch version {

	// version=docker passado no ldflags
	case "docker":
		println("\u001b[92mdocker\u001b[0m")
		fs := http.FileServer(http.FS(res))
		http.Handle("/manual/", http.StripPrefix("/", fs))
		println("\u001b[93mRun Server:", HOST_DOCKER, "\u001b[0m")

	// Não é necessário passar no ldflags main.version
	default:
		println("\u001b[93mdefault\u001b[0m")
		http.Handle("/", http.StripPrefix("/",
			http.FileServer(http.Dir("./manual"))))
		err := util.OpenBrowser(HOST)
		if err != nil {
			log.Println(err)
			return
		}
		println("\u001b[93mRun Server:", HOST, "\u001b[0m")
	}

	// Listen Server Run
	http.ListenAndServe(HTTP_PORT, nil)
}
