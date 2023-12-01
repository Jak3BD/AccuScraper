package locationsearch

type Response struct {
	Key           string `json:"key"`
	LocalizedName string `json:"localizedName"`

	AdministrativeArea struct {
		LocalizedName string `json:"localizedName"`
	} `json:"administrativeArea"`

	Country struct {
		LocalizedName string `json:"localizedName"`
	} `json:"country"`
}

type LocationSearch struct {
	Country            string `json:"country"`
	AdministrativeArea string `json:"administrativeArea"`
	LocalizedName      string `json:"localizedName"`
	Key                string `json:"key"`
}
