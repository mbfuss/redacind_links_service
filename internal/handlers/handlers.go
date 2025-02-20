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
	mux.HandleFunc("/", h.Distributor)
}

// Distributor - основной обработчик, распределяющий запросы.
func (h *Handler) Distributor(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		h.GoReduceLink(res, req)
	}
	if req.Method == http.MethodPost {
		h.AbbreviateLinks(res, req)
	}
}
