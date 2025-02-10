package memory_provider

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
	return p.ma.GetValueMemory(memoryKeyId)
}
func (p *provider) SetValueMemory(key string, value string) {
	p.ma.SetValueMemory(key, value)
}
