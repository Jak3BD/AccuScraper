package currentweather

import (
	"accuscraper/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
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

	url := fmt.Sprintf("https://www.accuweather.com/%s/%s/%s/%s/current-weather/%s", language, resolvedKey.Country, resolvedKey.LocalizedName, resolvedKey.ZIP, key)
	doc, err := utils.RequestDocument(language, url)
	if err != nil {
		utils.SendError(w, "Error getting document", err, http.StatusInternalServerError)
		return
	}

	var data CurrentWeather

	data.Date = utils.CleanerTxt(doc.Find(".subnav-pagination div").Text())
	data.Time = utils.CleanerTxt(doc.Find(".sub").Text())
	data.Temp = utils.CleanerTxt(doc.Find(".temp").Text())
	data.TempFeel = strings.ReplaceAll(utils.CleanerTxt(doc.Find(".current-weather-extra").First().Text()), "RealFeel® ", "")
	data.Phrase = utils.CleanerTxt(doc.Find(".phrase").First().Text())

	doc.Find(".detail-item").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 1:
			data.Wind = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		case 2:
			data.WindGusts = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		case 3:
			data.Humidity = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		case 4:
			data.DewPoint = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		case 5:
			data.AirPressure = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		case 6:
			data.Cloudiness = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		case 7:
			data.SightDistance = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		case 8:
			data.CeilingClouds = utils.CleanerTxt(s.Find("div").Eq(1).Text())
		}
	})

	icon, err := utils.GetSVG(doc.Find(".current-weather-info svg"))
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

	utils.SendJSON(w, data)
}
