package main

import (
	"fmt"
	"groupie/internal/config"
	"groupie/internal/handlers"
	"groupie/internal/render"
	"groupie/internal/repo"
	"groupie/internal/service"
	"net/http"
)

const (
	urlArtist   = "https://groupietrackers.herokuapp.com/api/artists"
	urlLocation = "https://groupietrackers.herokuapp.com/api/locations"
	urlRelation = "https://groupietrackers.herokuapp.com/api/relation"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("ERROR CREATING CACHE")
		return
	}
	app.TemplateCache = tc
	app.UseCache = true
	render.NewTemplate(&app)

	r := repo.New(urlArtist, urlRelation, urlLocation)
	s := service.New(r)
	h := handlers.New(s)

	mux := http.NewServeMux()

	mux.Handle("/static/", h.Handler)

	fmt.Println("Listening on http://localhost:8080/ ... ")
	mux.HandleFunc("/", h.IndexHandler)
	mux.HandleFunc("/artist", h.ArtistHandler)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Could not start sever on port http://localhost:8080/")
		return
	}
}
