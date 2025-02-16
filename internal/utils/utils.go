package utils

import (
	"fmt"
	"net/http"
)

func VerificationRequest(method string, ct string, res http.ResponseWriter, req *http.Request) error {
	if req.Method != method {
		return fmt.Errorf("не поддерживаемый http метод: %v", req.Method)
	}
	if req.Header.Get("Content-Type") != ct {
		http.Error(res, "Не поддерживаемый Content-Type", http.StatusUnsupportedMediaType)
		return fmt.Errorf("не поддерживаемый Content-Type: %v", req.Header.Get("Content-Type"))
	}
	return nil
}
