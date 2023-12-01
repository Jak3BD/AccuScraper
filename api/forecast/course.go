package forecast

import (
	"accuscraper/utils"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func course(doc *goquery.Document) Course {
	card := doc.Find(".weather-card")

	var data Course

	card.Each(func(i int, s *goquery.Selection) {
		qa, exists := s.Attr("data-qa")
		if exists {
			switch qa {
			case "todayWeatherCard":
				data.Today = getCourse(s)
			case "tonightWeatherCard":
				data.Tonight = getCourse(s)
			case "tomorrowWeatherCard":
				data.Tomorrow = getCourse(s)
			}
		}
	})

	return data
}

func getCourse(s *goquery.Selection) CourseData {
	var data CourseData

	data.Date = utils.CleanerTxt(s.Find(".sub-title").Text())
	data.Temp = utils.CleanerTxt(s.Find(".temp").Text())
	data.TempFeel = strings.ReplaceAll(utils.CleanerTxt(s.Find(".real-feel").Text()), "RealFeel® ", "")
	data.Phrase = utils.CleanerTxt(s.Find(".phrase").Text())

	icon, err := utils.GetSVG(s.Find("svg"))
	if err != nil {
		fmt.Println("Error getting SVG:", err)
	} else {
		data.Icon = icon
	}

	return data
}
