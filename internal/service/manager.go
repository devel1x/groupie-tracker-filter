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

func New(r repo.RepoI) ServiceI {
	return &service{
		r,
	}

}
