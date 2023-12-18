package service

import (
	"fmt"
	"groupie/internal/models"
	"strconv"
)

func (s *service) Artists() ([]models.Artist, error) {
	fmt.Print("Got all artists")
	artist, err := s.RepoI.Artists()
	if err != nil {

	}
	loc, err := s.RepoI.AllLoc()
	if err != nil {

	}
	for key, _ := range artist {
		artist[key].Location = &loc[key]
	}
	return artist, nil
}

func (s *service) Artist(id string) (*models.Artist, error) {
	if id == "" {
		return nil, fmt.Errorf(models.Errors[400].Msg)
	}
	if id[0] == '0' {
		return nil, fmt.Errorf(models.Errors[400].Msg)
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf(models.Errors[400].Msg)
	}

	if idInt < 1 {
		return nil, fmt.Errorf(models.Errors[404].Msg)
	}

	artist, err := s.RepoI.Artist(id)
	if err != nil {
		return nil, fmt.Errorf(models.Errors[500].Msg)
	}
	if artist.Members == nil {
		return nil, fmt.Errorf(models.Errors[404].Msg)
	}

	location, err := s.RepoI.Loc(id)
	if err != nil {
		return nil, fmt.Errorf(models.Errors[500].Msg)
	}
	artist.Location = location

	relation, err := s.RepoI.Rel(id)
	if err != nil {
		return nil, fmt.Errorf(models.Errors[500].Msg)
	}
	artist.Relation = relation
	return artist, nil
}
