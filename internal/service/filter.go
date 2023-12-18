package service

import (
	"groupie/internal/models"
	"strconv"
)

var err error

func (f *filter) Filter() ([]models.Artist, error) {
	if f.aDate != nil {
		f.artist, err = f.fAlbum()
		if err != nil {
			return nil, err
		}
	}
	if f.cDate != nil {
		f.artist, err = f.fAlbum()
		if err != nil {
			return nil, err
		}
	}
	if f.members != "" {
		f.artist, err = f.fAlbum()
		if err != nil {
			return nil, err
		}
	}
	if f.loc != "" {
		f.artist, err = f.fAlbum()
		if err != nil {
			return nil, err
		}
	}
	return f.artist, nil
}

func (f *filter) fCreation() ([]models.Artist, error) {
	var artist []models.Artist

	dates := make(map[string]int)
	for key, value := range f.cDate {
		dates[key], err = f.parseInt(value)
		if err != nil {
			return nil, err
		}
	}
	for key, value := range f.artist {
		if value.CreationDate >= dates["from"] && value.CreationDate <= dates["to"] {
			artist = append(artist, f.artist[key])
		}
	}
	return artist, nil
}

func (f *filter) fAlbum() ([]models.Artist, error) {
	// var artist []models.Artist

	// dates := make(map[string]int)
	// for key, value := range f.aDate {
	// 	dates[key], err = f.parseInt(value)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// for key, value := range f.artist {
	// 	if value.FirstAlbum >= dates["from"] && value.FirstAlbum <= dates["to"] {
	// 		artist = append(artist, f.artist[key])
	// 	}
	// }
	// return artist, nil
}

func (f *filter) fMembers() ([]models.Artist, error) {
	var artist []models.Artist
	
	n,err:=f.parseInt(f.members)
	if err!=nil{
		return nil, err
	}

	for key, value := range f.artist {
		var count int
		for range value.Members {
			count++
			}
		if count==n {
			artist = append(artist, f.artist[key])
		}
	}
	return artist, nil
}

func (f *filter) fLocation() ([]models.Artist, error) {
	var artist []models.Artist

	for key, value := range f.artist {
		var count int
		for range value.Location.Location {
			
			}
		if count==n {
			artist = append(artist, f.artist[key])
		}
	}
	return artist, nil
}

func (f *filter) parseInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
