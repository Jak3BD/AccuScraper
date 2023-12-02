# AccuScraper

AccuWeather scraping server for a self-hosted JSON API

The server uses the website [https://www.accuweather.com](https://www.accuweather.com) to scrape data and make it accessible as a json api

The documentation can be viewed in `openapi.yml` or interactively in the browser under the main route

## Usage
* `docker build -t accuscraper .`
* `docker run --name accuscraper -p 80:8080 -d accuscraper`

## Environment Variables
```bash
# defines the format in which the log is output
# options: pretty | json
# default: json
LOG_FORMAT

# defines when an event is logged
# options: 0-4 (0: debug, 1: info, 2: warn, 3: error, 4: fatal)
# default: 0
LOG_LEVEL
```