package handlers

import (
	"encoding/json"
	"groopie/fetchers"
	"net/http"
	"strings"
)

func JSONSearchPage(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("query"))
	// Выполняем поиск артистов
	artists, err := fetchers.FetchArtist()
	if err != nil {
		http.Error(w, "Ошибка при выполнении поиска", http.StatusInternalServerError)
		return
	}

	// Возвращаем результаты поиска в формате JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(artists); err != nil {
		http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
		return
	}
}
