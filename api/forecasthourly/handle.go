package forecasthourly

import (
	"accuscraper/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	url := fmt.Sprintf("https://www.accuweather.com/%s/%s/%s/%s/hourly-weather-forecast/%s", language, resolvedKey.Country, resolvedKey.LocalizedName, resolvedKey.ZIP, key)
	doc, err := utils.RequestDocument(language, url)
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
			fmt.Println("Error getting SVG:", err)
		} else {
			data.Icon = icon
		}

		forecast = append(forecast, data)
	})

	utils.SendJSON(w, forecast)
}
