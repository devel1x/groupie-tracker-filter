package repo

import "groupie/internal/models"

type RepoI interface {
	Artists() ([]models.Artist, error)
	Artist(string) (*models.Artist, error)
	Rel(n string) (*models.Relation, error)
	Loc(n string) (*models.Location, error)
}

type repo struct {
	urlArtists string
	urlRelation string
	urlLocation string
}

func New(url1, url2, url3 string) RepoI {
	return &repo{
		url1,
		url2,
		url3,
	}
}
