# AccuScraper

AccuWeather scraping server for a self-hosted JSON API

The server uses the website [https://www.accuweather.com](https://www.accuweather.com) to scrape data and make it accessible as a json api

The documentation can be viewed in `openapi.yml` or interactively in the browser under the main route

* `docker build -t accuscraper .`
* `docker run --name accuscraper -p 80:8080 -d accuscraper`