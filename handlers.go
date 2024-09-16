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
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, artists)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError, "internal server error")
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
	}
	t, err := template.ParseFiles(`templates/details.html`)
	if err != nil {
		ErrorHandler(w, req, http.StatusNotFound, "details.html not found")
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
	}
}
