package handlers

import (
	"groupie/internal/service"
	"net/http"
)

type handler struct {
	service.ServiceI
	http.Handler
}

func New(s service.ServiceI) *handler {
	return &handler{
		s,
		StaticHandler,
	}
}
