package router

import (
	"net/http"

	"ots/internal/api/handler"
)

func New(h *handler.Handler) http.Handler {
	mux := http.NewServeMux()
	
	// mux.HandleFunc("/shorten", h.Shorten)

	// mux.HandleFunc("/getoriginal", h.GetOriginal)

	return mux
}