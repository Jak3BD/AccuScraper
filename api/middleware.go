package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "origin, x-requested-with")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// for additional stuff, e.g. authentication
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Fields(map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.Path,
			"remote": r.RemoteAddr,
		}).Msg("Request received")

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
