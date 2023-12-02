package main

import (
	"accuscraper/api"
	"embed"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed index.html
var index embed.FS

//go:embed openapi.yml
var openapi embed.FS

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if format := os.Getenv("LOG_FORMAT"); format == "pretty" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	levelEnv := os.Getenv("LOG_LEVEL")
	if levelEnv == "" {
		return
	}

	level, err := strconv.Atoi(levelEnv)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse LOG_LEVEL")
	}

	switch level {
	case 0:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case 1:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case 2:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case 3:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case 4:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		log.Fatal().Msg("invalid log level")
	}
}

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

	log.Info().Msg("starting server")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal().Err(err).Msg("server stopped")
}
