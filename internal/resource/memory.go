package resource

// Memory структура для хранения данных о сайтах.
type Memory struct {
	// Urls - хэш-таблица для хранения информации о сайте:
	// k - сокращенная ссылка, v - сайт для редиректа.
	Urls map[string]string
}
