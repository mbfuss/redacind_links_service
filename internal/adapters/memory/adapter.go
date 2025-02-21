package memory

import (
	"1_increment_http_server/internal/resource"
	"fmt"
	"log"
)

type memoryAdapter struct {
	memory resource.Memory
}

func New() *memoryAdapter {
	return &memoryAdapter{
		memory: resource.Memory{
			Urls: make(map[string]string),
		},
	}
}

// GetValueMemory - получает полную ссылку для перехода.
func (ma *memoryAdapter) GetValueMemory(memoryKeyId string) (string, error) {
	// Проверка наличия ключа в мапе.
	value, exists := ma.memory.Urls[memoryKeyId]
	if !exists {
		return "", fmt.Errorf("данный ключ %s не найден", memoryKeyId)
	}
	return value, nil
}

// SetValueMemory - добавляет в хэш-таблицу информацию о сайте.
func (ma *memoryAdapter) SetValueMemory(key string, value string) {
	ma.memory.Urls[key] = value
	log.Println("Key:", key, "| OriginURL:", value)
}
