package memory_provider

import "fmt"

type memoryProvider interface {
	GetValueMemory(memoryKeyId string) (string, error)
	SetValueMemory(key string, value string)
}

// memoryProvider - провайдер для взаимодействия с memoryAdapter.
type provider struct {
	ma memoryProvider
}

func New(ma memoryProvider) *provider {
	return &provider{
		ma: ma,
	}
}

func (p *provider) GetValueMemory(memoryKeyId string) (string, error) {
	value, err := p.ma.GetValueMemory(memoryKeyId)
	if err != nil {
		return "", fmt.Errorf("не удалось найти ключ: %w", err)
	}
	return value, nil
}
func (p *provider) SetValueMemory(key string, value string) {
	p.ma.SetValueMemory(key, value)
}
