package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     *Location
	ConcertDates string `json:"concertDates"`
	Relation     *Relation
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
}

type LocationResponse struct {
	Index []Location `json:"index"`
}
