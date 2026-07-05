package handler

import (
	_ "encoding/json"
	_ "net/http"

	"ots/internal/service"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

var req struct {
    URL string `json:"url"`
}
