package handlers

import (
	"encoding/json"
	"groopie/fetchers"
	"groopie/models"
	"groopie/pkg"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func SearchPage(w http.ResponseWriter, r *http.Request) {

	query := strings.ToLower(r.URL.Query().Get("query"))

	var searchResult models.SearchResult

	locations, err := fetchers.FetchLocations()
	artists, err := fetchers.FetchArtist()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	for _, val := range locations.Index {
		for _, loc := range val.Locations {
			if strings.Contains(strings.ToLower(loc), query) {
				searchResult.Locations = append(searchResult.Locations, models.Locationsss{
					ID:        val.ID,
					Locations: val.Locations,
					Name:      artists[val.ID-1].Name,
				})
			}
		}
	}

	flag := true
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
	for _, val := range artists {
		if strings.Contains(strings.ToLower(val.Name), query) {
			searchResult.Artists = append(searchResult.Artists, val)
		}
	}

	searchResult1 := pkg.RemoveDublicates(searchResult)

	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {

		w.Header().Set("Content-Type", "application/json")
		encodedData, err := json.Marshal(searchResult1)
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
			err = tmpl.Execute(w, searchResult1)
		}
	}
}
