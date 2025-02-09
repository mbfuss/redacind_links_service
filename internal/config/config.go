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
	configPath string = "../../cfg/conf.env"
)

//Config – это структура, которая содержит настройки (например, Port).
//*Config – это указатель на структуру Config, то есть адрес объекта в памяти.
//&Config{Port: "8080"} – создаёт объект Config и возвращает его адрес.

// Load - возвращает структуру со значениями из конфиг файла.
func Load() *Config {
	err := godotenv.Load(configPath)
	if err != nil {
		fmt.Println("Error loading .env file")
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
