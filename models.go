package main

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// -------------------- Locations --------------------

type Locations struct {
	Index []LocationItem `json:"index"`
}

type LocationItem struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// -------------------- Dates --------------------

type Dates struct {
	Index []DateItem `json:"index"`
}

type DateItem struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// -------------------- Relations --------------------

type Relations struct {
	Index []RelationItem `json:"index"`
}

type RelationItem struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
