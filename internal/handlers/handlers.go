package handlers

import (
	"groupie/internal/models"
	"groupie/internal/render"
	"net/http"
	"strings"
	"text/template"
)

var StaticHandler = http.StripPrefix("/static/", preventDirListing(http.FileServer(http.Dir("./static"))))

var t *template.Template

func (h *handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, Errors[404])
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, Errors[405])
		return
	}

	Artists, err := h.ServiceI.Artists()
	// fmt.Print(Artists)
	if err != nil {
		ErrorHandler(w, r, Errors[500])
		return
	}

	Data := make(map[string]interface{})

	Data["Artists"] = Artists

	render.RenderTemplate(w, "index.page.html", &models.TemplateData{
		Data: Data,
	})
}

func preventDirListing(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || len(r.URL.Path) == 0 {
			ErrorHandler(w, r, Errors[404])
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func (h *handler) ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorHandler(w, r, Errors[404])
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, Errors[405])
		return
	}
	artistID := r.URL.Query().Get("id")

	Artist, err := h.ServiceI.Artist(artistID)
	if err != nil {
		if err.Error() == "Not found" {
			ErrorHandler(w, r, Errors[404])
			return
		}
		ErrorHandler(w, r, Errors[500])
		return
	}
	Data := make(map[string]interface{})

	Data["Artist"] = Artist

	render.RenderTemplate(w, "artist.page.html", &models.TemplateData{
		Data: Data,
	})
}
