package currentweather

type CurrentWeather struct {
	Date          string `json:"date"`
	Time          string `json:"time"`
	Temp          string `json:"temp"`
	TempFeel      string `json:"tempFeel"`
	Phrase        string `json:"phrase"`
	Wind          string `json:"wind"`
	WindGusts     string `json:"windGusts"`
	Humidity      string `json:"humidity"`
	DewPoint      string `json:"dewPoint"`
	AirPressure   string `json:"airPressure"`
	Cloudiness    string `json:"cloudiness"`
	SightDistance string `json:"sightDistance"`
	CeilingClouds string `json:"ceilingClouds"`
	Icon          string `json:"icon,omitempty"`
}
