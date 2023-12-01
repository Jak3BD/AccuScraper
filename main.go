package main

import (
	"accuscraper/api"
	"embed"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed index.html
var index embed.FS

//go:embed openapi.yml
var openapi embed.FS

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(index)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/openapi.yml", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(openapi)).ServeHTTP(w, r)
	})

	apiRoutes := mux.PathPrefix("/api").Subrouter()
	apiRoutes.Use(api.CorsMiddleware, api.Middleware)
	api.RoutesV1(apiRoutes.PathPrefix("/v1").Subrouter())

	err := http.ListenAndServe(":8080", mux)
	panic(err)
}
