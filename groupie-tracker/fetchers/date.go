package fetchers

import (
	"encoding/json"
	"groopie/models"
	"io"
	"net/http"
)

const datesAPI = "https://groupietrackers.herokuapp.com/api/dates"

func FetchDates() (models.Dates, error) {
	resp, err := http.Get(datesAPI)
	if err != nil {
		return models.Dates{}, err
	}
	defer resp.Body.Close()

	dataJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Dates{}, err
	}

	var dates models.Dates
	err = json.Unmarshal(dataJson, &dates)
	if err != nil {
		return models.Dates{}, err
	}
	return dates, nil

}
