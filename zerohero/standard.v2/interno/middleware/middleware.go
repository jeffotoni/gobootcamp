// JWT
// CORS
// LOGGER
// RATE LIMIT
// CACHE
// INSTRUMENTACAO
// BANCO DE DADOS
// GERAÇÃO DE LOGS SAÍDA

package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logger() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() {
				log.Printf(
					"\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m \033[1;44m%s \033[0m%s \033[1;44m%s\033[0m",
					r.Method,
					r.RequestURI,
					r.RemoteAddr,
					r.Header.Get("User-Agent"),
					r.URL.Path,
					time.Since(start))
			}()
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Use applies middlewares to a http.HandlerFunc
func Use(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
