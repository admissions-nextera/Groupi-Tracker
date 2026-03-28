package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ArtistCache []Artist

func InitCache() error {
	artists, err := getArtists()
	if err != nil {
		return err
	}
	ArtistCache = artists
	return nil
}

func getArtist(id int) (*Artist, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Request Field: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected Status: %d", resp.StatusCode)
	}
	var artist Artist
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return nil, fmt.Errorf("Field Decode : %w", err)
	}
	return &artist, nil
}

func getArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, fmt.Errorf("Request Field: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected Status: %d", resp.StatusCode)
	}
	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, fmt.Errorf("Field Decode : %w", err)
	}
	return artists, nil
}

func getRelations() (*Relations, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, fmt.Errorf("Request Field: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected Status: %d", resp.StatusCode)
	}
	var relations Relations
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, fmt.Errorf("Field Decode : %w", err)
	}
	return &relations, nil
}
