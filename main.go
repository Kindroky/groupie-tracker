package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	artists := fetchArtists()
	location := fetchLocation()
	dates := fetchDates()
	relation := fetchRelation()
	serverCreate(artists, relation, location, dates)
}

func serverCreate(artists []artistsStruc, relation []relationStruct, location []locationStruct, dates []datesStruct) {
	indexHandler := func(w http.ResponseWriter, req *http.Request) {
		IndexHandler(w, req, artists)
	}
	detailsHandler := func(w http.ResponseWriter, req *http.Request) {
		DetailsHandler(w, req, artists, relation, location, dates)
	}
	aboutHandler := func(w http.ResponseWriter, req *http.Request) {
		AboutHandler(w, req)
	}
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/details/", detailsHandler)
	mux.HandleFunc("/about", aboutHandler)

	server := &http.Server{
		Addr:              ":8080",          //adresse du server (le port choisi est à titre d'exemple)
		Handler:           mux,              // listes des handlers
		ReadHeaderTimeout: 10 * time.Second, // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second, // temps maximum d'écriture de la réponse
		IdleTimeout:       15 * time.Second, // temps maximum entre deux rêquetes
		MaxHeaderBytes:    1 << 20,          // 1 MB // maximum de bytes que le serveur va lire
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
