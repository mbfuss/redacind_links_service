package config

import (
	"1_increment_http_server/internal/resource"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ResourceConfig resource.Config
}

var (
	configPath = "cfg/conf.env"
)

//Config – это структура, которая содержит настройки (например, Port).
//*Config – это указатель на структуру Config, то есть адрес объекта в памяти.
//&Config{Port: "8080"} – создаёт объект Config и возвращает его адрес.

// Load - возвращает структуру со значениями из конфиг файла.
func Load() *Config {
	cwd, _ := os.Getwd()
	fmt.Println("Текущая директория:", cwd)
	err := godotenv.Load(configPath)
	if err != nil {
		fmt.Printf("Ошибка чтения конфигурационного файла: %v\n", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Будет установлен порт по умолчанию: 8080")
		port = "8080"
	}
	return &Config{
		ResourceConfig: resource.Config{
			Port: port,
		},
	}
}
