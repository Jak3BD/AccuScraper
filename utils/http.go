package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJSON(w http.ResponseWriter, fields interface{}) {
	err := json.NewEncoder(w).Encode(fields)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, msg string, err error, statusCode int) {
	log.Println(msg, err)

	w.WriteHeader(statusCode)
	SendJSON(w, map[string]interface{}{
		"error": msg,
	})
}
