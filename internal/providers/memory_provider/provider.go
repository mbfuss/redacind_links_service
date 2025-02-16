package memory_provider

import "fmt"

type memoryAdapter interface {
	GetValueMemory(memoryKeyId string) (string, error)
	SetValueMemory(key string, value string)
}

// MemoryProvider - провайдер для взаимодействия с memoryAdapter.
type MemoryProvider struct {
	ma memoryAdapter
}

func New(ma memoryAdapter) *MemoryProvider {
	return &MemoryProvider{
		ma: ma,
	}
}

func (p *MemoryProvider) GetValueMemory(memoryKeyId string) (string, error) {
	value, err := p.ma.GetValueMemory(memoryKeyId)
	if err != nil {
		return "", fmt.Errorf("не удалось найти ключ: %w", err)
	}
	return value, nil
}

func (p *MemoryProvider) SetValueMemory(key string, value string) {
	p.ma.SetValueMemory(key, value)
}
