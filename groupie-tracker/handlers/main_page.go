package handlers

import (
	"groopie/fetchers"
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		tmpl, err := template.ParseFiles("views/404.html")
		if err != nil {
			http.Error(w, "Ошибка при парсинга шаблона", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusNotFound)
		err = tmpl.Execute(w, nil)
		return
	}

	artists, err := fetchers.FetchArtist()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("views/artists.html")
	if err != nil {
		http.Error(w, "Ошибка при парсинга шаблона", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "Ошибка при отображении данных", http.StatusInternalServerError)
		return
	}
}
