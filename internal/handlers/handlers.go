package handlers

import (
	"1_increment_http_server/internal/business/shortener"
	"1_increment_http_server/internal/providers/memory_provider"
	"net/http"
)

type Handler struct {
	sh shortener.Shortener
	mp memory_provider.MemoryProvider
}

func New(sh shortener.Shortener, mp memory_provider.MemoryProvider) *Handler {
	return &Handler{
		sh: sh,
		mp: mp,
	}
}

// RegisterHandlers - регистрации обработчиков HTTP-запросов.
func (h *Handler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Distributor)
}

// Distributor - обработчик, распределяет запросы для /.
func (h *Handler) Distributor(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		h.GoReduceLink(res, req)
	}
	if req.Method == http.MethodPost {
		h.AbbreviateLinks(res, req)
	}
}
