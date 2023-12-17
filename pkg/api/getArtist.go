package api

import (
	"encoding/json"
	"fmt"
	"groupie/pkg/models"
	"net/http"
	"time"
)

var client *http.Client

const (
	urlArtist   = "https://groupietrackers.herokuapp.com/api/artists"
	urlLocation = "https://groupietrackers.herokuapp.com/api/locations"
	urlRelation = "https://groupietrackers.herokuapp.com/api/relation"
)

func init() {
	client = &http.Client{Timeout: 10 * time.Second}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetArtists() (*[]models.Artist, error) {
	var artists []models.Artist
	err := GetJson(urlArtist, &artists)
	if err != nil {
		return nil, err
	}

	return &artists, nil
}

func GetArtistByID(id string) (*models.Artist, error) {
	var artist models.Artist
	urlId := urlArtist + "/" + id
	fmt.Print(urlId)
	err := GetJson(urlId, &artist)
	if err != nil {
		return nil, err
	}
	if artist.Members == nil {
		
		return &artist, fmt.Errorf("Not found")
	}

	location, err := GetLocationByID(id)
	if err != nil {
		return nil, err
	}
	artist.Location = location

	relation, err := GetRelationByID(id)
	if err != nil {
		return nil, err
	}
	artist.Relation = relation

	return &artist, nil
}

func GetLocationByID(id string) (*models.Location, error) {
	url := urlLocation + "/" + id
	var location models.Location
	err := GetJson(url, &location)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &location, nil
}

func GetRelationByID(id string) (*models.Relation, error) {
	url := urlRelation + "/" + id
	var relation models.Relation
	err := GetJson(url, &relation)
	if err != nil {
		return nil, err
	}

	return &relation, nil
}
