package api

import (
	"accuscraper/api/currentweather"
	"accuscraper/api/forecast"
	"accuscraper/api/forecastdaily"
	"accuscraper/api/forecasthourly"
	"accuscraper/api/locationsearch"

	"github.com/gorilla/mux"
)

func RoutesV1(r *mux.Router) {
	r.HandleFunc("/location-search", locationsearch.Handle).Methods("GET")
	r.HandleFunc("/forecast", forecast.Handle).Methods("GET")
	r.HandleFunc("/forecast/hourly", forecasthourly.Handle).Methods("GET")
	r.HandleFunc("/forecast/daily", forecastdaily.Handle).Methods("GET")
	r.HandleFunc("/current-weather", currentweather.Handle).Methods("GET")
}
