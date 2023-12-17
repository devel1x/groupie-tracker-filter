package handlers

import (
	"groupie/pkg/models"
	"groupie/pkg/render"
	"net/http"
)

type errs struct {
	Code int
	Msg  string
}

var (
	errors = map[int]errs{
		404: {
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound),
		},
		405: {
			http.StatusMethodNotAllowed,
			http.StatusText(http.StatusMethodNotAllowed),
		},
		400: {
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		},
		500: {
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		},
	}
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, myErr errs) {
	w.WriteHeader(myErr.Code)

	render.RenderTemplate(w, "error.page.html", &models.TemplateData{
		Error:   myErr.Code,
		Warning: myErr.Msg,
	})

}
