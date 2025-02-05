package main

import (
	"1_increment_http_server/internal/config"
	"fmt"
	"net/http"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	fmt.Printf("Запущен сервер на порту %v", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		fmt.Println("Запуск сервера: %w", err)
	}
}
