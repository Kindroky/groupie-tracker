package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// function that handles the requests related to the errors (404, 500)
func ErrorHandler(w http.ResponseWriter, req *http.Request, errorCode int, errorMessage string) {
	w.WriteHeader(errorCode) //display error message in the terminal of the navigator

	t, err := template.New(`error.html`).ParseFiles(`templates/error.html`) // parse through the files to find the error file
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
	err = t.Execute(w, errData) // execute the error html file
	if err != nil {
		fmt.Println(err)
		return
	}
}

// function to handle the index requests
func IndexHandler(w http.ResponseWriter, req *http.Request, artists []artistsStruc) {
	if req.URL.Path != "/" { // if url is not root, display 404 error
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

// function to handle the requests related to the details page
func DetailsHandler(w http.ResponseWriter, req *http.Request, artists []artistsStruc, relation []relationStruct, location []locationStruct, dates []datesStruct) {
	artistID := req.URL.Query().Get("ID") // get the id of the chosen artist
	var artistIDint int                   //initiate variable artistIDint to store the id
	if artistID == "" {
		ErrorHandler(w, req, http.StatusInternalServerError, "no ID in URL")
		return
	}
	artistIDint, err := strconv.Atoi(artistID) //convert the id to int and store it in the proper variable
	if err != nil {
		fmt.Println(err)
		return
	}
	if artistIDint > len(artists) || artistIDint < 1 { //handle the case where the id is not in the range of artists displayed
		ErrorHandler(w, req, http.StatusNotFound, "Page not found")
		return
	}
	t, err := template.ParseFiles(`templates/details.html`)
	if err != nil {
		ErrorHandler(w, req, http.StatusNotFound, "details.html not found")
		return
	}
	data := pageData{ //get the corresponding api values, depending on the id we fetched
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

// function that handles requests related to the about page
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
