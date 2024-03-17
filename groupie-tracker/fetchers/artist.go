package fetchers

import (
	"encoding/json"
	"groopie/models"
	"io"
	"net/http"
)

const artistAPI = "https://groupietrackers.herokuapp.com/api/artists"

func FetchArtist() ([]models.Artist, error) {

	resp, err := http.Get(artistAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	dataJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return []models.Artist{}, err
	}
	err = json.Unmarshal(dataJson, &artists)

	return artists, nil
}
