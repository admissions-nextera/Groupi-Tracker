package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	HomeTemplate   *template.Template
	ArtistTemplate *template.Template
)

func init() {
	HomeTemplate = template.Must(template.ParseFiles("templates/index.html"))
	ArtistTemplate = template.Must(template.ParseFiles("templates/artist.html"))
}

func HomeHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(resp, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(resp, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := HomeTemplate.Execute(resp, ArtistCache); err != nil {
		log.Println("Template error:", err)
	}
}

func ArtistHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(resp, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		http.Error(resp, "Bad Request", http.StatusBadRequest)
		return
	}

	artist, err := getArtist(id)
	if err != nil {
		http.Error(resp, "Not Found", http.StatusNotFound)
		return
	}

	rel, err := getRelations()
	if err != nil {
		log.Println(err)
		http.Error(resp, "Internal Server Error", 500)
		return
	}

	var relation RelationItem
	for _, item := range rel.Index {
		if item.ID == id {
			relation = item
			break
		}
	}

	data := struct {
		Artist   *Artist
		Relation RelationItem
	}{
		Artist:   artist,
		Relation: relation,
	}

	if err := ArtistTemplate.Execute(resp, data); err != nil {
		log.Println(err)
		http.Error(resp, "Internal Server Error", 500)
		return
	}
}

func SearchHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(resp, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	query := strings.ToLower(req.URL.Query().Get("q"))

	var results []Artist
	for _, a := range ArtistCache {
		// Check Artist Name
		nameMatch := strings.Contains(strings.ToLower(a.Name), query)

		// Check all Members
		memberMatch := false
		for _, m := range a.Members {
			if strings.Contains(strings.ToLower(m), query) {
				memberMatch = true
				break
			}
		}

		// Check Creation Date
		dateMatch := strings.Contains(strconv.Itoa(a.CreationDate), query)

		if nameMatch || memberMatch || dateMatch {
			results = append(results, a)
		}
	}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(results)
}
