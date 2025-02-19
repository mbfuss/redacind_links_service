package handlers

import (
	"1_increment_http_server/internal/utils"
	"fmt"
	"net/http"
)

func (h *Handler) GoReduceLink(res http.ResponseWriter, req *http.Request) {
	err := utils.VerificationRequest(http.MethodGet, "text/plain", res, req)
	if err != nil {
		http.Error(res, fmt.Sprintf("Произошла ошибка: %v", err), http.StatusBadRequest)
		return
	}
	linkSite := req.URL.Path
	res.Write([]byte(linkSite))
}
