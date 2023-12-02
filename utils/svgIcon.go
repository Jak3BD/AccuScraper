package utils

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

func GetSVG(icon *goquery.Selection) (string, error) {
	viewBox, _ := icon.Attr("viewBox")
	width, _ := icon.Attr("width")
	height, _ := icon.Attr("height")

	// Remove color from SVG
	// icon.Find("path").Each(func(i int, s *goquery.Selection) {
	// 	s.RemoveAttr("stroke")
	// })

	iconHtml, err := icon.Html()
	if err != nil {
		return "", err
	}

	svg := fmt.Sprintf(`<svg viewBox="%s" width="%s" height="%s">%s</svg>`, viewBox, width, height, iconHtml)

	tmpl, err := template.New("sanitize").Parse(svg)
	if err != nil {
		return "", err
	}

	var sanitizedSVG strings.Builder
	err = tmpl.Execute(&sanitizedSVG, nil)
	if err != nil {
		return "", err
	}

	return sanitizedSVG.String(), nil
}
