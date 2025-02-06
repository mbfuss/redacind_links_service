package main

import (
	"1_increment_http_server/internal/config"
	"1_increment_http_server/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	handler.RegisterHandlers(mux)

	fmt.Printf("Запущен сервер на порту %v", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		fmt.Println("Запуск сервера: %w", err)
	}
}
