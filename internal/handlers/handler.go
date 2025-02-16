package handlers

import (
	"1_increment_http_server/internal/service/shortener"
	"net/http"
)

type Handler struct {
	sh shortener.Shortener
}

func New(sh shortener.Shortener) *Handler {
	return &Handler{
		sh: sh,
	}
}

// RegisterHandlers - регистрации обработчиков HTTP-запросов.
func (h *Handler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", h.AbbreviateLinks)
	mux.HandleFunc("/{id}", h.GoReduceLink)
}
