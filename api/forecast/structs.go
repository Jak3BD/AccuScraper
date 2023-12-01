package forecast

type Forecast struct {
	Warning        *Warning       `json:"warning,omitempty"`
	CurrentWeather CurrentWeather `json:"currentWeather"`
	AirQuality     AirQuality     `json:"airQuality"`
	Course         Course         `json:"course"`
}

type Warning struct {
	Count  string `json:"count"`
	Phrase string `json:"phrase"`
}

type CurrentWeather struct {
	Time       string `json:"time"`
	Temp       string `json:"temp"`
	TempFeel   string `json:"tempFeel"`
	Phrase     string `json:"phrase"`
	AirQuality string `json:"airQuality"`
	Wind       string `json:"wind"`
	WindGusts  string `json:"windGusts"`
	Icon       string `json:"icon,omitempty"`
}

type AirQuality struct {
	Date   string `json:"date"`
	Value  string `json:"value"`
	Unit   string `json:"unit"`
	Phrase string `json:"phrase"`
}

type Course struct {
	Today    CourseData `json:"today"`
	Tonight  CourseData `json:"tonight"`
	Tomorrow CourseData `json:"tomorrow"`
}

type CourseData struct {
	Date     string `json:"date"`
	Temp     string `json:"temp"`
	TempFeel string `json:"tempFeel"`
	Phrase   string `json:"phrase"`
	Icon     string `json:"icon,omitempty"`
}
