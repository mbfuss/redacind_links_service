package main

import (
	"1_increment_http_server/internal/adapters/memory"
	"1_increment_http_server/internal/config"
	"1_increment_http_server/internal/handlers"
	"1_increment_http_server/internal/providers/memory_provider"
	"1_increment_http_server/internal/service/shortener"
	"fmt"
	"log"
	"net/http"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Произошла паника в приложении: %s", err)
		}
	}()
	cfg := config.Load()

	mux := http.NewServeMux()

	h := handlers.New(*shortener.New(*memory_provider.New(memory.New())))
	h.RegisterHandlers(mux)

	fmt.Println("Запущен сервер на порту:", cfg.ResourceConfig.Port)
	err := http.ListenAndServe(":"+cfg.ResourceConfig.Port, mux)
	if err != nil {
		fmt.Println("Запуск сервера: %w", err)
	}
}
