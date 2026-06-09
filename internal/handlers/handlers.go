package handlers

import (
	"html/template"
	"net/http"

	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/render"
	"ascii-art-web/internal/validation"
)

type PageData struct {
	Text   string
	Banner string
	Error  string
	Result string
}

func renderHomePage(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	renderHomePage(w, PageData{})
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	err = validation.ValidateText(text)
	if err != nil {
		renderHomePage(w, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  err.Error(),
		})
		return
	}

	err = validation.ValidateBanner(bannerName)
	if err != nil {
		renderHomePage(w, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  err.Error(),
		})
		return
	}

	bannerFile := bannerName + ".txt"

	bannerData, err := banner.Load(bannerFile)
	if err != nil {
		renderHomePage(w, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  err.Error(),
		})
		return
	}

	result := render.Generate(text, bannerData)

	renderHomePage(w, PageData{
		Text:   text,
		Banner: bannerName,
		Result: result,
	})
}