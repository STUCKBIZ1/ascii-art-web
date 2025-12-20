package main

import (
	"html/template"
	"net/http"
	"os"

	ascii_art "ascii-art-web/ascii-art/src"
)

type PageData struct {
	Result string
	Error  string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	data := PageData{
		Result: "",
	}

	tmpl.Execute(w, data)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	bannerData, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}


	font := ascii_art.Sep_Fonts(string(bannerData))
	result := ascii_art.Chars_To_Art(font, text)

	data := PageData{
		Result: result,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	http.ListenAndServe(":8080", nil)
}
