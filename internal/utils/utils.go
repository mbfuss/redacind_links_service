package utils

import (
	"fmt"
	"net/http"
)

func VerificationRequest(ct string, res http.ResponseWriter, req *http.Request) error {
	if req.Header.Get("Content-Type") != ct {
		http.Error(res, "Не поддерживаемый Content-Type", http.StatusUnsupportedMediaType)
		return fmt.Errorf("не поддерживаемый Content-Type: %v", req.Header.Get("Content-Type"))
	}
	return nil
}
