package handler

import (
	srvc "github.com/Communinst/GolangWebStore/backend/service"
)

type Handler struct {
	service *srvc.Service
}

func New(service *srvc.Service) *Handler {
	return &Handler{
		service: service,
	}
}
