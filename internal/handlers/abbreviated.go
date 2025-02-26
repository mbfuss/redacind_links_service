package handlers

import (
	"1_increment_http_server/internal/utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AbbreviateLinks - отправляет оригинальный URL в хранилище и связывает его с укороченным.
func (h *Handler) AbbreviateLinks(res http.ResponseWriter, req *http.Request) {
	err := utils.VerificationRequest("text/plain", res, req)
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
		return
	}

	shorten := h.sh.ShortenJoin(originUrl)

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(h.cfg.ResourceConfig.Host + ":" + h.cfg.ResourceConfig.Port + "/" + shorten))
}
