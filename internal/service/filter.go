package service

import (
	"groupie/internal/models"
	"strconv"
	"strings"
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
	var artist []models.Artist

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
	return artist, nil
}

func (f *filter) fMembers() ([]models.Artist, error) {
	var artist []models.Artist

	n, err := f.parseInt(f.members)
	if err != nil {
		return nil, err
	}

	for key, value := range f.artist {
		var count int
		for range value.Members {
			count++
		}
		if count == n {
			artist = append(artist, f.artist[key])
		}
	}
	return artist, nil
}

func (f *filter) fLocation() ([]models.Artist, error) {
	var artist []models.Artist

	for key, value := range f.artist {
		flag := false
		for _, value := range value.Location.Location {
			if value == f.loc {
				flag = true
			}
		}
		if flag {
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
func GetLocF(artist []models.Artist) map[string][]string {
	var (
		loc = make(map[string][]string)
	)
	for _, value := range artist {
		for _, value := range value.Location.Location {
			a := strings.Split(value, "-")
			a[0] = strings.Title(strings.Replace(a[0], "_", "-", -1))
			a[1] = strings.Title(a[1])
			if _, exists := loc[a[1]]; !exists {
				loc[a[1]] = []string{}
				loc[a[1]] = append(loc[a[1]], a[0])
			} else {
				if !contains(loc[a[1]], a[0]) {
					loc[a[1]] = append(loc[a[1]], a[0])
				}
			}
		}
	}
	return loc
}

func contains(arr []string, a string) bool {
	for _, value := range arr {
		if value == a {
			return true
		}
	}
	return false
}
