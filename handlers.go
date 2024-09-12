package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	t, err := template.ParseFiles(`templates/index.html`)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, req, http.StatusNotFound, "index.html not found")
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, artists)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, req, http.StatusInternalServerError, "internal server error")
	}
}
