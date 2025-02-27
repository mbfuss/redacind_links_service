package handlers

import (
	"1_increment_http_server/internal/utils"
	"fmt"
	"net/http"
	"strings"
)

// GoReduceLink - получает оригинальный URL по укороченному.
func (h *Handler) GoReduceLink(res http.ResponseWriter, req *http.Request) {
	err := utils.VerificationRequest("text/plain", res, req)
	if err != nil {
		http.Error(res, fmt.Sprintf("Произошла ошибка: %v", err), http.StatusBadRequest)
		return
	}
	linkSite := strings.Split(req.URL.Path, "/")

	originURl, err := h.mp.GetValueMemory(strings.Join(linkSite, ""))
	if err != nil {
		http.Error(res, fmt.Sprintf("Произошла ошибка: %v", err), http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusTemporaryRedirect)
	res.Write([]byte(originURl))
}
