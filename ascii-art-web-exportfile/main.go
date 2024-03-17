package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"web/ascii-art"
)

var outputtxt string

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.Handle("/stats/", http.StripPrefix("/stats/", http.FileServer(http.Dir("stats"))))
	mux.HandleFunc("/ascii-art", onlyText)

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		//log.Fatal(err)
		return
	}

}

// Обработчик главной страницы
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound) //404
		http.ServeFile(w, r, "stats/404.html")
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) //405
		return
	}
	http.ServeFile(w, r, "stats/index.html")
}

func onlyText(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound) //404
		http.ServeFile(w, r, "404.html")
		return
	}
	switch r.Method {
	case http.MethodGet:
		getPerson(w)
	case http.MethodPost:
		postPerson(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed) //405
	}
}

func getPerson(w http.ResponseWriter) {
	fmt.Fprintf(w, "%v", outputtxt)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	Text := r.FormValue("fname")
	Style := r.FormValue("styleName")
	second := ""
	outputtxt, second = ascii.AsciiImpl(Text, Style)
	if second == "no Latin" {
		tmpl, _ := template.ParseFiles("stats/400.html")
		w.WriteHeader(http.StatusBadRequest) //400
		tmpl.Execute(w, nil)
		return
	} else if second == "no banner" {
		tmpl, _ := template.ParseFiles("stats/500.html")
		w.WriteHeader(http.StatusInternalServerError) //500
		tmpl.Execute(w, nil)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
