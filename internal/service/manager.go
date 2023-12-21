package service

import (
	"groupie/internal/models"
	"groupie/internal/repo"
)

type ServiceI interface {
	Artists() ([]models.Artist, error)
	Artist(id string) (*models.Artist, error)
}

type service struct {
	repo.RepoI
}

type FilterI interface {
	Filter() ([]models.Artist, error)
}

type filter struct {
	artist  []models.Artist
	aDate   string
	cDate   map[string]string
	members string
	loc     map[string]string
}

func New(r repo.RepoI) ServiceI {
	return &service{
		r,
	}
}

func NewFilter(artist []models.Artist, aDate, cDateF, cDateT, members string, loc map[string]string) FilterI {
	var cDate map[string]string
	if cDateF != "" || cDateT != "" {
		cDate = make(map[string]string)
		cDate["from"] = cDateF
		cDate["to"] = cDateT
	}
	return &filter{
		artist,
		aDate,
		cDate,
		members,
		loc,
	}
}
