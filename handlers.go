package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, errorCode int, errorMessage string) {
	w.WriteHeader(errorCode)

	t, err := template.New(`error.html`).ParseFiles(`templates/error.html`)
	if err != nil {
		fmt.Println(err)
		return
	}
	errData := struct {
		ErrorCode    int
		ErrorMessage string
	}{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, errData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func IndexHandler(w http.ResponseWriter, req *http.Request, artists []artistsStruc) {
	if req.URL.Path != "/" {
		ErrorHandler(w, req, http.StatusNotFound, "Page not found")
		return
	}
	t, err := template.ParseFiles(`templates/index.html`)
	if err != nil {
		ErrorHandler(w, req, http.StatusNotFound, "index.html not found")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, artists)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError, "internal server error")
		return
	}
}

func DetailsHandler(w http.ResponseWriter, req *http.Request, artists []artistsStruc, relation []relationStruct, location []locationStruct, dates []datesStruct) {
	artistID := req.URL.Query().Get("ID")
	var artistIDint int
	if artistID == "" {
		ErrorHandler(w, req, http.StatusInternalServerError, "no ID in URL")
		return
	}
	artistIDint, err := strconv.Atoi(artistID)
	if err != nil {
		fmt.Println(err)
		return
	}
	if artistIDint > len(artists) || artistIDint < 1 {
		ErrorHandler(w, req, http.StatusNotFound, "Page not found")
		return
	}
	t, err := template.ParseFiles(`templates/details.html`)
	if err != nil {
		ErrorHandler(w, req, http.StatusNotFound, "details.html not found")
		return
	}
	data := pageData{
		Artists:   artists[artistIDint-1],
		Locations: location[artistIDint-1],
		Dates:     dates[artistIDint-1],
		Relation:  relation[artistIDint-1],
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, req, http.StatusInternalServerError, "internal server error")
		return
	}
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles(`templates/about.html`)
	if err != nil {
		ErrorHandler(w, req, http.StatusNotFound, "about.html not found")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError, "internal server error")
		return
	}
}
