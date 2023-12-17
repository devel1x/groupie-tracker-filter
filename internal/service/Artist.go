package service

import (
	"fmt"
	"groupie/internal/models"
	"strconv"
	"groupie/internal/handlers"
)

func (s *service) Artists() ([]models.Artist, error) {
	return s.RepoI.Artists()
}

func (s *service) Artist(id string) (*models.Artist, error) {
	if id == "" {
		return nil, fmt.Errorf(handlers.Errors[400].Msg)
	}
	if id[0] == '0' {
		return nil, fmt.Errorf(handlers.Errors[400].Msg)
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf(handlers.Errors[400].Msg)
	}

	if idInt < 1 {
		return nil, fmt.Errorf(handlers.Errors[404].Msg)
	}

	artist, err := s.RepoI.Artist(id)

	location, err := s.RepoI.Loc(id)
	if err != nil {
		return nil, fmt.Errorf(handlers.Errors[500].Msg)
	}
	artist.Location = location

	relation, err := s.RepoI.Rel(id)
	if err != nil {
		return nil, fmt.Errorf(handlers.Errors[500].Msg)
	}
	artist.Relation = relation
	return artist, nil
}
