package locationsearch

import (
	"accuscraper/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	language := params.Get("language")
	query := params.Get("query")

	baseUrl := fmt.Sprintf("https://www.accuweather.com/web-api/autocomplete?language=%s&query=%s", language, query)

	req, err := utils.NewRequest(language, baseUrl)
	if err != nil {
		utils.SendError(w, "Error creating request", err, http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		utils.SendError(w, "Error sending request", err, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var data []Response
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		utils.SendError(w, "Error decoding response", err, http.StatusInternalServerError)
		return
	}

	var sendData []LocationSearch
	for _, v := range data {
		sendData = append(sendData, LocationSearch{
			Country:            v.Country.LocalizedName,
			AdministrativeArea: v.AdministrativeArea.LocalizedName,
			LocalizedName:      v.LocalizedName,
			Key:                v.Key,
		})
	}

	utils.SendJSON(w, sendData)
}
