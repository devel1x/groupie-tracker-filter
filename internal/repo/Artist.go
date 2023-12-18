package repo

import (
	"encoding/json"
	"fmt"
	"groupie/internal/models"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{Timeout: 10 * time.Second}
}

func (r *repo) GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func (r *repo) Artists() ([]models.Artist, error) {
	var artists []models.Artist
	err := r.GetJson(r.urlArtists, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (r *repo) Artist(id string) (*models.Artist, error) {
	var artist models.Artist
	urlId := r.urlArtists + "/" + id
	fmt.Print(urlId)
	err := r.GetJson(urlId, &artist)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (r *repo) Loc(id string) (*models.Location, error) {
	url := r.urlLocation + "/" + id
	var location models.Location
	err := r.GetJson(url, &location)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &location, nil
}

func (r *repo) AllLoc() ([]models.Location, error) {

	var location []models.LocationResponse

	err := r.GetJson(r.urlLocation, &location)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	var allLocations []models.Location
	for _, locationResponse := range location {
		allLocations = append(allLocations, locationResponse.Index...)
	}
	return allLocations, nil
}

func (r *repo) Rel(id string) (*models.Relation, error) {
	url := r.urlRelation + "/" + id
	var relation models.Relation
	err := r.GetJson(url, &relation)
	if err != nil {
		return nil, err
	}

	return &relation, nil
}
