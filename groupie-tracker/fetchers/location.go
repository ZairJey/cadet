package fetchers

import (
	"encoding/json"
	"groopie/models"
	"io"
	"net/http"
)

const locationsAPI = "https://groupietrackers.herokuapp.com/api/locations"

func FetchLocations() (models.Locations, error) {

	resp, err := http.Get(locationsAPI)
	if err != nil {
		return models.Locations{}, err
	}
	defer resp.Body.Close()

	dataJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Locations{}, err
	}

	var locations models.Locations
	err = json.Unmarshal(dataJson, &locations)
	if err != nil {
		return models.Locations{}, err
	}
	return locations, nil

}
