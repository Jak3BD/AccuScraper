package forecasthourly

import (
	"accuscraper/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := r.URL.Query()
	language := params.Get("language")
	key := params.Get("key")

	resolvedKey, err := utils.ResolveKey(ctx, language, key)
	if err != nil {
		utils.SendError(w, "Error resolving key", err, http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("https://www.accuweather.com/%s/%s/%s/%s/hourly-weather-forecast/%s", language, resolvedKey.Country, resolvedKey.LocalizedName, resolvedKey.ZIP, key)
	doc, err := utils.RequestDocument(ctx, language, url)
	if err != nil {
		utils.SendError(w, "Error getting document", err, http.StatusInternalServerError)
		return
	}

	var forecast []Forecast

	doc.Find(".accordion-item").Each(func(i int, s *goquery.Selection) {
		var data Forecast

		data.Time = utils.CleanerTxt(s.Find(".date").Text())
		data.Temp = utils.CleanerTxt(s.Find(".temp").Text())
		data.TempFeel = strings.ReplaceAll(utils.CleanerTxt(s.Find(".real-feel__text").Text()), "RealFeel® ", "")
		data.Phrase = utils.CleanerTxt(s.Find(".phrase").Text())
		data.Humidity = utils.CleanerTxt(s.Find(".precip").Text())

		icon, err := utils.GetSVG(s.Find("svg"))
		if err != nil {
			log.Error().Err(err).Fields(map[string]interface{}{
				"country":       resolvedKey.Country,
				"localizedName": resolvedKey.LocalizedName,
				"zip":           resolvedKey.ZIP,
				"key":           key,
			}).Msg("error getting SVG")
		} else {
			data.Icon = icon
		}

		forecast = append(forecast, data)
	})

	utils.SendJSON(w, forecast)
}
