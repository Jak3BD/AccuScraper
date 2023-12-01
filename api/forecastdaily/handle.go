package forecastdaily

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

	url := fmt.Sprintf("https://www.accuweather.com/%s/%s/%s/%s/daily-weather-forecast/%s", language, resolvedKey.Country, resolvedKey.LocalizedName, resolvedKey.ZIP, key)
	doc, err := utils.RequestDocument(language, url)
	if err != nil {
		utils.SendError(w, "Error getting document", err, http.StatusInternalServerError)
		return
	}

	var forecasts []Forecast

	doc.Find(".daily-wrapper").Each(func(i int, s *goquery.Selection) {
		var data Forecast

		data.Day = utils.CleanerTxt(s.Find(".date .dow").Text())
		data.Date = utils.CleanerTxt(s.Find(".date .sub").Text())
		data.TempHigh = utils.CleanerTxt(s.Find(".high").Text())
		data.TempLow = strings.ReplaceAll(utils.CleanerTxt(s.Find(".low").Text()), "/", "")
		data.Phrase = utils.CleanerTxt(s.Find(".phrase").Text())
		data.Humidity = utils.CleanerTxt(s.Find(".precip").Text())

		rightPanel := s.Find(".right .panel-item")
		data.UVIndex = utils.CleanerTxt(rightPanel.Eq(0).Find(".value").Text())
		data.Wind = utils.CleanerTxt(rightPanel.Eq(1).Find(".value").Text())

		icon, err := utils.GetSVG(s.Find("svg"))
		if err != nil {
			fmt.Println("Error getting SVG:", err)
		} else {
			data.Icon = icon
		}

		warnings := s.Find(".inline-alert-banners")
		if warnings.Length() > 0 {
			var warningsData []Warning
			warnings.Find(".inline-alert").Each(func(i int, s *goquery.Selection) {
				var warning Warning

				subheading := s.Find(".inline-alert-subheading p")
				warning.Description = utils.CleanerTxt(subheading.Eq(0).Text())
				warning.TimeSpan = utils.CleanerTxt(subheading.Eq(1).Text())

				warningsData = append(warningsData, warning)
			})

			data.Warnings = &warningsData
		}

		forecasts = append(forecasts, data)
	})

	utils.SendJSON(w, forecasts)
}
