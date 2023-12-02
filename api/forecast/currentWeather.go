package forecast

import (
	"accuscraper/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
)

func currentWeather(doc *goquery.Document) CurrentWeather {
	card := doc.Find(".cur-con-weather-card")

	var data CurrentWeather

	data.Time = utils.CleanerTxt(card.Find(".cur-con-weather-card__subtitle").Text())
	data.Temp = utils.CleanerTxt(card.Find(".temp").Text())
	data.TempFeel = strings.ReplaceAll(utils.CleanerTxt(card.Find(".real-feel").Text()), "RealFeel® ", "")
	data.Phrase = utils.CleanerTxt(card.Find(".phrase").Text())

	icon, err := utils.GetSVG(card.Find("svg"))
	if err != nil {
		log.Error().Err(err).Msg("error getting SVG")
	} else {
		data.Icon = icon
	}

	details := card.Find(".details-container .spaced-content")
	details.Each(func(i int, s *goquery.Selection) {
		airQuality := 0
		wind := 1
		windGusts := 2

		if details.Length() == 4 {
			airQuality = 1
			wind = 2
			windGusts = 3
		}

		switch i {
		case airQuality:
			data.AirQuality = utils.CleanerTxt(s.Find(".value").Text())
		case wind:
			data.Wind = utils.CleanerTxt(s.Find(".value").Text())
		case windGusts:
			data.WindGusts = utils.CleanerTxt(s.Find(".value").Text())
		}
	})

	return data
}
