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

type ErrorPageData struct {
	Code    int
	Message string
}

func renderHomePage(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderErrorPage(w, http.StatusNotFound, "Page Not Found")
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func renderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)

	err = tmpl.Execute(w, ErrorPageData{
		Code:    statusCode,
		Message: message,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		renderErrorPage(
			w,
			http.StatusNotFound,
			"Page Not Found",
		)
		return
	}

	if r.Method != http.MethodGet {
		renderErrorPage(
			w,
			http.StatusMethodNotAllowed,
			"Method Not Allowed",
		)
		return
	}

	renderHomePage(w, PageData{})
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusBadRequest, "Bad Request")
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
