package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// GoReduceLink - получает оригинальный URL по укороченному.
func (h *Handler) GoReduceLink(res http.ResponseWriter, req *http.Request) {
	linkSite := strings.Split(req.URL.Path, "/")

	originURl, err := h.mp.GetValueMemory(strings.Join(linkSite, ""))
	if err != nil {
		http.Error(res, fmt.Sprintf("Произошла ошибка: %v", err), http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", originURl)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
