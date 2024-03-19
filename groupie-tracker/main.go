package main

import (
	"groopie/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/artists/", handlers.ArtsPage)
	http.HandleFunc("/search", handlers.SearchPage)

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		return
	}
}
