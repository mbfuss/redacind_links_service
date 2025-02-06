package handler

import "net/http"

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", nil)
	mux.HandleFunc("/{id}", nil)
}
