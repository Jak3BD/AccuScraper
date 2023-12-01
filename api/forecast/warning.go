package forecast

import (
	"accuscraper/utils"

	"github.com/PuerkitoBio/goquery"
)

func warning(doc *goquery.Document) (Warning, bool) {
	card := doc.Find(".severe-alert-banner")
	if card.Length() == 0 {
		return Warning{}, false
	}

	var data Warning

	data.Count = utils.CleanerTxt(card.Find(".severe-alert-banner__count").Text())
	data.Phrase = utils.CleanerTxt(card.Find(".severe-alert-banner__type").Text())

	return data, true
}
