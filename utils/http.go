package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func SendJSON(w http.ResponseWriter, fields interface{}) {
	err := json.NewEncoder(w).Encode(fields)
	if err != nil {
		log.Error().Err(err).Msg("error encoding JSON")

		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, msg string, err error, statusCode int) {
	log.Error().Err(err).Msg(msg)

	w.WriteHeader(statusCode)
	SendJSON(w, map[string]interface{}{
		"error": msg,
	})
}
