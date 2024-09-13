package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type artistsStruc struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type relationNoIndex struct {
	Index []relationStruct `json:"index"`
}
type relationStruct struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func fetchArtists() []artistsStruc {
	var artists []artistsStruc
	apiArtist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer apiArtist.Body.Close()
	err = json.NewDecoder(apiArtist.Body).Decode(&artists)
	if err != nil {
		fmt.Println("error:", err)
	}
	return artists
}

func fetchRelation() []relationStruct {
	var relation relationNoIndex
	apiRelation, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}
	defer apiRelation.Body.Close()
	err = json.NewDecoder(apiRelation.Body).Decode(&relation)
	if err != nil {
		fmt.Println("error:", err)
	}
	return relation.Index
}
