package handlers

import (
	"groopie/fetchers"
	"groopie/models"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func ArtsPage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := path.Base(r.URL.Path)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Ошибка при получении ID артиста", http.StatusNotFound)
		return
	}
	artists, err := fetchers.FetchArtist()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}
	relations, err := fetchers.FetchRelation()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}
	locations, err := fetchers.FetchLocations()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}
	dates, err := fetchers.FetchDates()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	var artist models.Artist
	for _, v := range artists {
		if v.ID == idInt {
			artist = v
			break
		}
	}

	if artist.ID == 0 {
		http.Error(w, "Артист не найден", http.StatusNotFound)
		return
	}

	var relationInArtist models.Relationss
	for _, rel := range relations.Index {
		if rel.ID == idInt {
			relationInArtist = rel
			break
		}
	}
	var locationInArtist models.Locationss
	for _, loc := range locations.Index {
		if loc.ID == idInt {
			locationInArtist = loc
			break
		}
	}
	var datesInArtist models.Datess
	for _, dat := range dates.Index {
		if dat.ID == idInt {
			datesInArtist = dat
			break
		}
	}

	tmpl, err := template.ParseFiles("views/artist.html")
	if err != nil {
		http.Error(w, "Ошибка при парсинга шаблона", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct {
		Artist    models.Artist
		Locations models.Locationss
		Relations models.Relationss
		Dates     models.Datess
	}{
		Artist:    artist,
		Locations: locationInArtist,
		Relations: relationInArtist,
		Dates:     datesInArtist,
	})
	if err != nil {
		http.Error(w, "Ошибка при отображении данных", http.StatusInternalServerError)
	}
}
