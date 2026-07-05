package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"ots/internal/api/dto"
	"ots/internal/apperrors"
	"ots/internal/service"
)

type Handler struct {
	svc *service.Service
	log *slog.Logger
}

func New(svc *service.Service, log *slog.Logger) *Handler {
	return &Handler{svc: svc, log: log}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req dto.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request body"})
		return
	}
	defer r.Body.Close()

	if err := validateURL(req.URL); err != nil {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: "invalid url"})
		return
	}

	short, err := h.svc.Shorten(req.URL)
	if err != nil {
		h.respondError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, dto.ShortenResponse{ShortURL: short})
}

func (h *Handler) GetOriginal(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: "invalid short code"})
		return
	}

	original, err := h.svc.GetOriginal(code)
	if err != nil {
		h.respondError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.OriginalResponse{URL: original})
}

func (h *Handler) respondError(w http.ResponseWriter, r *http.Request, err error) {
	status, msg := mapError(err)
	if status >= http.StatusInternalServerError {
		h.log.Error("request failed",
			"err", err,
			"method", r.Method,
			"path", r.URL.Path,
		)
	}
	writeJSON(w, status, dto.ErrorResponse{Error: msg})
}

func mapError(err error) (int, string) {
	switch {
	case errors.Is(err, apperrors.ErrNotFound):
		return http.StatusNotFound, "link not found"
	case errors.Is(err, apperrors.ErrInvalidURL):
		return http.StatusBadRequest, "invalid url"
	case errors.Is(err, apperrors.ErrConflict):
		return http.StatusConflict, "conflict"
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}

func validateURL(raw string) error {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return apperrors.ErrInvalidURL
	}

	u, err := url.ParseRequestURI(raw)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return apperrors.ErrInvalidURL
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
