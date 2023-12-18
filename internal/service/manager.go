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
	aDate   map[string]string
	cDate   map[string]string
	members string
	loc     string
}

func New(r repo.RepoI) ServiceI {
	return &service{
		r,
	}

}

func NewFilter(artist []models.Artist, aDateF, aDateT, cDateF, cDateT, members, loc string) FilterI {
	var aDate map[string]string
	if aDateF!="" && aDateT!=""{
		aDate = make(map[string]string)
		aDate["from"] = aDateF
		aDate["to"] = aDateT
	}
	var cDate map[string]string
	if cDateF!="" && cDateT!=""{
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
