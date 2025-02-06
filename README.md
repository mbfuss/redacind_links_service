Инкремент 1

Задание по треку «Сервис сокращения URL»
Чтобы написать сервис, который будет сжимать длинные URL до нескольких символов, для начала вам нужно разработать сервер.
Сервер должен быть доступен по адресу http://localhost:8080 и предоставлять два эндпоинта:
Эндпоинт с методом POST и путём /. Сервер принимает в теле запроса строку URL как text/plain и возвращает ответ с кодом 201 и сокращённым URL как text/plain.
Пример запроса к серверу:
POST / HTTP/1.1
Host: localhost:8080
Content-Type: text/plain
https://practicum.yandex.ru/

Пример ответа от сервера:
HTTP/1.1 201 Created
Content-Type: text/plain
Content-Length: 30
http://localhost:8080/EwHXdJfB

Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB). В случае успешной обработки запроса сервер возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
Пример запроса к серверу:
GET /EwHXdJfB HTTP/1.1
Host: localhost:8080
Content-Type: text/plain
Пример ответа от сервера:
HTTP/1.1 307 Temporary Redirect
Location: https://practicum.yandex.ru/
На любой некорректный запрос сервер должен возвращать ответ с кодом 400.

Доп. информация о проекте:

Такая структура проекта основана на "Clean Architecture" (Чистая архитектура) и частично включает элементы "Hexagonal Architecture" (Гексагональная архитектура, aka Ports & Adapters).

Разбор паттерна:

Разделение на слои

cmd/ – точка входа, только инициализация сервера.

internal/handler/ – слой обработки HTTP-запросов (контроллеры).

internal/service/ – бизнес-логика, отделённая от HTTP.

internal/storage/ – слой работы с хранилищем данных.

url-shortener/
│── cmd/
│   └── server/
│       └── main.go            # Точка входа
│── internal/
│   ├── handler/
│   │   ├── shorten.go         # Обработчик сокращения URL
│   │   ├── redirect.go        # Обработчик редиректа
│   │   └── handler.go         # Регистрация обработчиков
│   ├── storage/
│   │   ├── memory.go          # Временное хранилище (карта)
│   │   ├── storage.go         # Интерфейс хранилища
│   └── service/
│       ├── shortener.go       # Логика сокращения URL
│── pkg/
│   ├── logger/
│   │   └── logger.go          # Логирование
│── config/
│   └── config.go              # Конфигурация (порт, БД и т. д.)
│── go.mod                     # Модуль Go
│── go.sum                     # Зависимости
│── README.md                  # Документация