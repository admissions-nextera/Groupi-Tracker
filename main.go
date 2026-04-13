package main

import (
	"log"
	"net/http"
)

func main() {

	if err := InitCache(); err != nil {
		log.Fatalf("Could not initialize data: %v", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/search", SearchHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
