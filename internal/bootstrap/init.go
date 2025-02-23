package bootstrap

import (
	"1_increment_http_server/internal/adapters/memory"
	"1_increment_http_server/internal/business/shortener"
	"1_increment_http_server/internal/config"
	"1_increment_http_server/internal/handlers"
	"1_increment_http_server/internal/providers/memory_provider"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	cfg *config.Config
	mux *http.ServeMux
}

func (a *App) Initialize() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatalf("Произошла паника в приложении: %s\n", err)
		}
	}()

	a.mux = http.NewServeMux()
	ma := memory.New()
	mp := memory_provider.New(ma)
	sh := shortener.New(*mp)
	h := handlers.New(*sh, *mp)
	h.RegisterHandlers(a.mux)
}

func (a *App) StartServer() error {
	cfg := a.cfg.Load()
	fmt.Println("Запущен сервер на порту:", cfg.ResourceConfig.Port)
	err := http.ListenAndServe(":"+cfg.ResourceConfig.Port, a.mux)
	if err != nil {
		return fmt.Errorf("Запуск сервера: %v\n", err)
	}
	return nil
}
