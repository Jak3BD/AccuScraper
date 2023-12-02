package utils

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ResolvedKey struct {
	Country       string
	LocalizedName string
	ZIP           string
}

var resolvedKeys = make(map[string]ResolvedKey)

func NewRequest(ctx context.Context, language, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	req.Header.Set("Accept-Language", language)

	return req, nil
}

func ResolveKey(ctx context.Context, language, key string) (ResolvedKey, error) {
	if resolved, ok := resolvedKeys[key]; ok {
		return resolved, nil
	}

	req, err := NewRequest(ctx, language, "https://www.accuweather.com/web-api/three-day-redirect?key="+key)
	if err != nil {
		return ResolvedKey{}, err
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Do(req)
	if err != nil {
		return ResolvedKey{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 302 {
		return ResolvedKey{}, errors.New("Invalid status code")
	}

	// Location: /en/us/los-angeles/90012/weather-forecast/347625
	location := res.Header.Get("Location")
	if location == "" {
		return ResolvedKey{}, errors.New("No location header")
	}

	var resolved ResolvedKey

	segments := strings.Split(location, "/")
	if len(segments) < 6 {
		return ResolvedKey{}, errors.New("Invalid location header")
	}

	resolved.Country = segments[2]
	resolved.LocalizedName = segments[3]
	resolved.ZIP = segments[4]

	resolvedKeys[key] = resolved
	return resolved, nil
}

func RequestDocument(ctx context.Context, language, url string) (*goquery.Document, error) {
	req, err := NewRequest(ctx, language, url)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid status code")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
