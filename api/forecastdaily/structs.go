package forecastdaily

type Forecast struct {
	Day      string     `json:"day"`
	Date     string     `json:"date"`
	TempHigh string     `json:"tempHigh"`
	TempLow  string     `json:"tempLow"`
	Phrase   string     `json:"phrase"`
	Humidity string     `json:"humidity"`
	UVIndex  string     `json:"uvIndex"`
	Wind     string     `json:"wind"`
	Warnings *[]Warning `json:"warnings,omitempty"`
	Icon     string     `json:"icon,omitempty"`
}

type Warning struct {
	Description string `json:"description"`
	TimeSpan    string `json:"timeSpan"`
}
