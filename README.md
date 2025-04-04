Инструкция для запуска

Открываем консоль

1. docker-compose up -d
2. go run .\cmd\main.go


# 📚 Books REST API

## 🔧 Технологии

- Go
- Gin (веб-фреймворк)
- GORM (ORM для работы с базой данных)
- SQLite / PostgreSQL / MySQL (можно выбрать)


📘 API эндпоинты
Базовый путь: /api/v1/books

Метод	URL	Описание
GET	/	Получить список всех книг
POST	/	Создать новую книгу
GET	/:id	Получить книгу по ID
PUT	/:id	Обновить информацию о книге
DELETE	/:id	Удалить книгу по ID