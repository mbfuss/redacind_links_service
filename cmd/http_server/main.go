package main

import (
	"1_increment_http_server/internal/adapters/memory"
	"1_increment_http_server/internal/providers/memory_provider"
	"fmt"
	"log"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Произошла паника в приложении: %s", err)
		}
	}()

	test := memory_provider.New(memory.New())
	fmt.Println(test.GetValueMemory("1"))
	test.SetValueMemory("1", "123")
	fmt.Println(test.GetValueMemory("1"))

	//cfg := config.Load()
	//
	//mux := http.NewServeMux()
	//handler.RegisterHandlers(mux)
	//
	//fmt.Printf("Запущен сервер на порту %v", cfg.ResourceConfig.Port)
	//err := http.ListenAndServe(":"+cfg.ResourceConfig.Port, mux)
	//if err != nil {
	//	fmt.Println("Запуск сервера: %w", err)
	//}
}
