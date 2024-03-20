package pkg

import "groopie/models"

func RemoveDublicates(stru models.SearchResult) models.SearchResult {
	uniqueArtists := make(map[int]models.Artist)
	for _, val := range stru.Artists {
		uniqueArtists[val.ID] = val
	}

	var uniqueArtistsList []models.Artist
	for _, val := range uniqueArtists {
		uniqueArtistsList = append(uniqueArtistsList, val)
	}

	uniqueLocations := make(map[int]models.Locationsss)
	for _, location := range stru.Locations {
		if _, found := uniqueLocations[location.ID]; !found {
			uniqueLocations[location.ID] = models.Locationsss{
				ID:        location.ID,
				Locations: location.Locations,
				Name:      location.Name,
			}
		}
	}
	var uniqueLocationsList []models.Locationsss
	for _, location := range uniqueLocations {
		uniqueLocationsList = append(uniqueLocationsList, location)
	}

	return models.SearchResult{
		Artists:   uniqueArtistsList,
		Locations: uniqueLocationsList,
	}
}
