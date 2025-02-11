package handler

import "net/http"

type Handler struct {
}

func (h *Handler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", nil)
	mux.HandleFunc("/{id}", nil)
}
