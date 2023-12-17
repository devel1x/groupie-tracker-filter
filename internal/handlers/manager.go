package handlers

import (
	"groupie/internal/service"
	"net/http"
	"time"
)

type handler struct {
	service.ServiceI
	http.Handler
	useCash      bool
	useCashTimer *time.Timer
}

func New(s service.ServiceI) *handler {
	return &handler{
		s,
		StaticHandler,
		false,
		nil,
	}
}
