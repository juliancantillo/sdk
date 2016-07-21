package sdk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Movie struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	Rated string `json:"Rated"`
	Released string `json:"Released"`
	Runtime string `json:"Runtime"`
	Genre string `json:"Genre"`
	Director string `json:"Director"`
	Writer string `json:"Writer"`
	Actors string `json:"Actors"`
	Plot string `json:"Plot"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Awards string `json:"Awards"`
	Poster string `json:"Poster"`
	Metascore string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes string `json:"imdbVotes"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	Response string `json:"Response"`
}

func NewMovie(imdbID string) *Movie {
	// QueryEscape escapes the movie id so
	// it can be safely placed inside a URL query
	safeImdbID := url.QueryEscape(imdbID)

	url := fmt.Sprintf("http://www.omdbapi.com/?i=%s&plot=short&r=json", safeImdbID)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a clien
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var movie Movie

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		log.Println(err)
	}

	return &movie
}