package forecast

import (
	"accuscraper/utils"

	"github.com/PuerkitoBio/goquery"
)

func airQuality(doc *goquery.Document) AirQuality {
	card := doc.Find(".air-quality-card")

	var data AirQuality

	data.Date = utils.CleanerTxt(card.Find(".date").Text())
	data.Value = utils.CleanerTxt(card.Find(".aq-number").Text())
	data.Unit = utils.CleanerTxt(card.Find(".aq-unit").Text())
	data.Phrase = utils.CleanerTxt(card.Find(".category-text").Text())

	return data
}
