package router

import (
	"net/http"

	"ots/internal/api/handler"
)

func New(h *handler.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", h.Shorten)
	mux.HandleFunc("GET /{code}", h.GetOriginal)

	return mux
}
