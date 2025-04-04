package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// All structures needed for scraping the APIs, NoIndexes ones are to avoid index in json
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
type apiNoIndex1 struct {
	IndexRel []relationStruct `json:"index"`
}
type relationStruct struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type apiNoIndex2 struct {
	IndexLoc []locationStruct `json:"index"`
}
type locationStruct struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type apiNoIndex3 struct {
	IndexDat []datesStruct `json:"index"`
}
type datesStruct struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Structure combining all 4 APIs
type pageData struct {
	Artists   artistsStruc
	Locations locationStruct
	Dates     datesStruct
	Relation  relationStruct
}

func fetchArtists() []artistsStruc {
	var artists []artistsStruc
	apiArtist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists") // Get function to scrap
	if err != nil {
		log.Fatal(err)
	}
	defer apiArtist.Body.Close()
	err = json.NewDecoder(apiArtist.Body).Decode(&artists) // Decoding json file from API
	if err != nil {
		fmt.Println("error:", err)
	}
	return artists
}

func fetchRelation() []relationStruct {
	var relation apiNoIndex1
	apiRelation, err := http.Get("https://groupietrackers.herokuapp.com/api/relation") // Get function to scrap
	if err != nil {
		log.Fatal(err)
	}
	defer apiRelation.Body.Close()
	err = json.NewDecoder(apiRelation.Body).Decode(&relation) // Decoding json file from API
	if err != nil {
		fmt.Println("error:", err)
	}
	return relation.IndexRel
}

func fetchLocation() []locationStruct {
	var locations apiNoIndex2
	apiLocations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations") // Get function to scrap
	if err != nil {
		log.Fatal(err)
	}
	defer apiLocations.Body.Close()
	err = json.NewDecoder(apiLocations.Body).Decode(&locations) // Decoding json file from API
	if err != nil {
		fmt.Println("error:", err)
	}
	return locations.IndexLoc
}

func fetchDates() []datesStruct {
	var dates apiNoIndex3
	apiDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates") // Get function to scrap
	if err != nil {
		log.Fatal(err)
	}
	defer apiDates.Body.Close()
	err = json.NewDecoder(apiDates.Body).Decode(&dates) // Decoding json file from API
	if err != nil {
		fmt.Println("error:", err)
	}
	return dates.IndexDat
}
