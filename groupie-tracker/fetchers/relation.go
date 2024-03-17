package fetchers

import (
	"encoding/json"
	"groopie/models"
	"io"
	"net/http"
)

const relationAPI = "https://groupietrackers.herokuapp.com/api/relation"

func FetchRelation() (models.Relations, error) {

	resp, err := http.Get(relationAPI)
	if err != nil {
		return models.Relations{}, err
	}
	defer resp.Body.Close()

	var relations models.Relations

	dataJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Relations{}, err
	}
	err = json.Unmarshal(dataJson, &relations)
	return relations, nil
}
