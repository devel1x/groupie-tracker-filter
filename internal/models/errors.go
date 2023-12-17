package models

import "net/http"

type Errs struct {
	Code int
	Msg  string
}

var (
	Errors = map[int]Errs{
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
