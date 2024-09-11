package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type artistStruc struct {
	id           int      `json:"id"`
	name         string   `json:"name"`
	image        string   `json:"image"`
	members      []string `json:"members"`
	creationDate int      `json:"creationDate"`
	firstAlbum   string   `json:"firstAlbum"`
	locations    string   `json:"locations"`
	concertDates string   `json:"concertDates"`
	relations    string   `json:"relations"`
}

func main() {

	var artists []artistStruc

	apiArtist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(apiArtist.Body)
	apiArtist.Body.Close()
	if apiArtist.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", apiArtist.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	/* 	var data map[string]interface{}
	 */erro := json.Unmarshal([]byte(body), &artists)
	if erro != nil {
		fmt.Println("error:", erro)
	}
	fmt.Printf("%+v", artists)

	/* mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", TBD)

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
	} */

	/* res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body) */
}
