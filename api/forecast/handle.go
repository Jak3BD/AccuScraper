package forecast

import (
	"accuscraper/utils"
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	language := params.Get("language")
	key := params.Get("key")

	resolvedKey, err := utils.ResolveKey(language, key)
	if err != nil {
		utils.SendError(w, "Error resolving key", err, http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("https://www.accuweather.com/%s/%s/%s/%s/weather-forecast/%s", language, resolvedKey.Country, resolvedKey.LocalizedName, resolvedKey.ZIP, key)
	doc, err := utils.RequestDocument(language, url)
	if err != nil {
		utils.SendError(w, "Error getting document", err, http.StatusInternalServerError)
		return
	}

	var forecast Forecast

	warning, exists := warning(doc)
	if exists {
		forecast.Warning = &warning
	}

	forecast.CurrentWeather = currentWeather(doc)
	forecast.AirQuality = airQuality(doc)
	forecast.Course = course(doc)

	utils.SendJSON(w, forecast)
}
