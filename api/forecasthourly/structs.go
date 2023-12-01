package forecasthourly

type Forecast struct {
	Time     string `json:"time"`
	Temp     string `json:"temp"`
	TempFeel string `json:"tempFeel"`
	Phrase   string `json:"phrase"`
	Humidity string `json:"humidity"`
	Icon     string `json:"icon,omitempty"`
}
