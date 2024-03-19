package handlers

import (
	"encoding/json"
	"groopie/fetchers"
	"groopie/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type SeachResult struct {
	Artists   []models.Artist     `json:"artists"`
	Locations []models.Locationss `json:"locations"`
}

func SearchPage(w http.ResponseWriter, r *http.Request) {

	query := strings.ToLower(r.URL.Query().Get("query"))

	var searchResult SeachResult

	locations, err := fetchers.FetchLocations()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	//var idInLocations []int
	for _, val := range locations.Index {
		for _, loc := range val.Locations {
			if strings.Contains(strings.ToLower(loc), query) {
				searchResult.Locations = append(searchResult.Locations, models.Locationss{
					ID:        val.ID,
					Locations: val.Locations,
				})
				//idInLocations = append(idInLocations, val.ID)
			}
		}
	}

	artists, err := fetchers.FetchArtist()
	if err != nil {
		return
	}
	flag := true
	//var ArtistsResult []models.Artist
	for _, v := range artists {
		if strings.Contains(strings.ToLower(v.Name), query) ||
			strings.Contains(strings.ToLower(v.FirstAlbum), query) ||
			strings.Contains(strings.ToLower(strconv.Itoa(v.CreationDate)), query) {
			searchResult.Artists = append(searchResult.Artists, v)
			flag = false
		} else {
			flag = true
		}
		for _, val := range v.Members {
			if strings.Contains(strings.ToLower(val), query) && flag {
				searchResult.Artists = append(searchResult.Artists, v)
				break
			}
		}
	}
	//for _, member := range artists {
	//	for _, val := range member.Members {
	//		if strings.Contains(strings.ToLower(val), query) {
	//			ArtistsResult = append(ArtistsResult, member)
	//			break
	//		}
	//	}
	//}

	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {

		w.Header().Set("Content-Type", "application/json")
		encodedData, err := json.Marshal(searchResult)
		if err != nil {
			http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
			return
		}
		//fmt.Println(string(encodedData))
		_, err = w.Write(encodedData)
		if err != nil {
			http.Error(w, "Ошибка при записи данных в ответ", http.StatusInternalServerError)
			return
		}
	} else {
		if len(searchResult.Artists) == 0 && len(searchResult.Locations) == 0 {
			tmpl, err := template.ParseFiles("views/404.html")
			if err != nil {
				return
			}
			err = tmpl.Execute(w, nil)

		} else {
			tmpl, err := template.ParseFiles("views/artists.html")
			if err != nil {
				return
			}
			err = tmpl.Execute(w, searchResult)
		}
	}
}

//func IfConnect(idIwki []int, id int) bool {
//	for _, num := range idIwki {
//		if id == num {
//			return true
//		}
//	}
//	return false
//}
