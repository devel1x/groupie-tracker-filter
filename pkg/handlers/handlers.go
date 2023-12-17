package handlers

import (
	"groupie/pkg/api"
	"groupie/pkg/models"
	"groupie/pkg/render"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var t *template.Template

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, errors[404])
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, errors[405])
		return
	}

	Artists, err := api.GetArtists()
	// fmt.Print(Artists)
	if err != nil {
		ErrorHandler(w, r, errors[500])
		return
	}

	Data := make(map[string]interface{})

	Data["Artists"] = Artists

	render.RenderTemplate(w, "index.page.html", &models.TemplateData{
		Data: Data,
	})
}

var StaticHandler = http.StripPrefix("/static/", preventDirListing(http.FileServer(http.Dir("./static"))))

func preventDirListing(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || len(r.URL.Path) == 0 {
			ErrorHandler(w, r, errors[404])
			return
		}
		h.ServeHTTP(w, r)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorHandler(w, r, errors[404])
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, errors[405])
		return
	}
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		ErrorHandler(w, r, errors[400])
		return
	}
	idInt, err := strconv.Atoi(artistID)
	if artistID[0] == '0' {
		ErrorHandler(w, r, errors[400])
		return
	}
	if err != nil {
		ErrorHandler(w, r, errors[400])
		return
	}
	if idInt < 1 {
		ErrorHandler(w, r, errors[404])
		return
	}

	Artist, err := api.GetArtistByID(artistID)
	// fmt.Print(Artist)
	if err != nil {
		if err.Error() == "Not found" {
			ErrorHandler(w, r, errors[404])
			return
		}
		ErrorHandler(w, r, errors[500])
		return
	}
	Data := make(map[string]interface{})

	Data["Artist"] = Artist

	render.RenderTemplate(w, "artist.page.html", &models.TemplateData{
		Data: Data,
	})
}
