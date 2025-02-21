package handlers

import (
	"1_increment_http_server/internal/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (h *Handler) AbbreviateLinks(res http.ResponseWriter, req *http.Request) {
	err := utils.VerificationRequest(http.MethodPost, "text/plain", res, req)
	if err != nil {
		http.Error(res, fmt.Sprintf("Произошла ошибка: %v", err), http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}

	originUrl := string(body)
	if originUrl == "" {
		http.Error(res, "Тело запроса не может быть пустым", http.StatusBadRequest)
	}
	h.sh.ShortenJoin(originUrl)
}
