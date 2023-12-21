package service

import (
	"fmt"
	"groupie/internal/models"
	"strconv"
	"strings"
)

var err error

func (f *filter) Filter() ([]models.Artist, error) {
	if f.aDate != "" {
		fmt.Println("aDate")
		f.artist, err = f.fAlbum()
		if err != nil {
			return nil, err
		}
	}
	if f.cDate != nil {
		fmt.Println("cDate")
		f.artist, err = f.fCreation()
		if err != nil {
			return nil, err
		}
	}
	if f.members != "" {
		fmt.Println("members")
		f.artist, err = f.fMembers()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(f.loc)
	if f.loc != nil {
		fmt.Println("location")
		f.artist, err = f.fLocation()
		if err != nil {
			return nil, err
		}
	}

	return f.artist, nil
}

func (f *filter) fCreation() ([]models.Artist, error) {
	var artist []models.Artist
	if f.cDate["from"] == "" {
		f.cDate["from"] = "1900"
	}
	if f.cDate["to"] == "" {
		f.cDate["to"] = "2023"
	}
	dates := make(map[string]int)
	for key, value := range f.cDate {
		dates[key], err = f.parseInt(value)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(f.cDate["from"])
	fmt.Println(f.cDate["to"])
	for key, value := range f.artist {
		if value.CreationDate >= dates["from"] && value.CreationDate <= dates["to"] {
			artist = append(artist, f.artist[key])
		}
	}

	return artist, nil
}

func (f *filter) fAlbum() ([]models.Artist, error) {
	var artist []models.Artist

	fDate, err := f.parseInt(f.aDate)
	if err != nil {
		return nil, err
	}

	for key, value := range f.artist {
		date, err := f.parseInt(value.FirstAlbum[6:])
		if err != nil {
			return nil, err
		}
		if date == fDate {
			artist = append(artist, f.artist[key])
		}
	}
	return artist, nil
}

// todo
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
	for _, valueA := range f.artist {
		a := valueA.Location
		for key, value := range a.CityMap {
			if _, exists := f.loc[key]; exists {
				if f.loc[key] == "" {
					artist = append(artist, valueA)
					continue
				} else if contains(value, f.loc[key]) {
					artist = append(artist, valueA)
					continue
				}
			}
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

func GetLocF(artist []models.Artist) ([]models.Artist, map[string][]string) {
	loc := make(map[string][]string)
	for _, valueA := range artist {
		valueA.Location.CityMap = make(map[string][]string)
		for _, value := range valueA.Location.Location {
			a := strings.Split(value, "-")
			a[0] = strings.Title(strings.Replace(a[0], "_", "-", -1))
			a[1] = strings.Title(strings.Replace(a[1], "_", " ", -1))
			if _, exists := valueA.Location.CityMap[a[1]]; !exists {
				valueA.Location.CityMap[a[1]] = []string{}
				valueA.Location.CityMap[a[1]] = append(valueA.Location.CityMap[a[1]], a[0])

			} else {
				if !contains(valueA.Location.CityMap[a[1]], a[0]) {
					valueA.Location.CityMap[a[1]] = append(valueA.Location.CityMap[a[1]], a[0])
				}
			}
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

	return artist, loc
}

func contains(arr []string, a string) bool {
	for _, value := range arr {
		if value == a {
			return true
		}
	}
	return false
}
