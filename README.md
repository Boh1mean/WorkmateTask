## I/O bound задача

Данный сервис — это минималистичное REST API-приложение на Go для имитации и управления длительными I/O-bound задачами. Каждая задача имеет случайную продолжительность выполнения и проходит через состояния: `pending` → `running` → `completed`.

---

## Обзор

Проект создан для демонстрации:

- работы с in-memory хранилищем,
- имитации длительных задач с помощью `goroutine` и `time.Sleep`,
- обработки задач через сервисный слой (`service`),
- построения REST API с использованием `net/http`,
- чистой архитектуры (handler → service → storage),
- покрытия логики unit-тестами.

---

## Возможности

- Создание задачи
- Получение информации о задаче по ID
- Удаление задачи
- Отслеживание статуса выполнения (`pending`, `running`, `completed`)

---
## Структура проекта 

```
workmateTask/
├── cmd/
│   └── main.go           # Точка входа
├── internal/
│   ├── model/            # Структура Task
│   ├── service/          # Бизнес-логика и in-memory хранилище
│   └── transport/        # HTTP handlers и роутинг
├── go.mod

```
---

## Требования

- Go 1.21+
- Bash или совместимый терминал для запуска команд

---

## Установка и запуск

```bash
# Клонируем репозиторий
git clone https://github.com/Boh1mean/workmateTask.git
cd workmateTask

# Установка зависимостей
go mod tidy

# Запуск приложения
go run cmd/main.go
```
Адрес сервера:  http://localhost:8080

## Service-Endpoints
| Метод | Путь               | Действие               |
|-------|--------------------|------------------------|
| POST  | /task              | Создать новую задачу   |
| GET   | /task?id={uuid}    | Получить статус задачи |
| DELETE| /task?id={uuid}    | Удалить задачу         |

## Формат ответа JSON
```json
{
  "id": "b35fc8ba-1f04-4f6a-b623-b346f63f7d09",
  "status": "running",
  "created_at": "2025-06-26T14:30:00Z",
  "duration": "4m15s"
}
```
* `id`: Уникальный идентификатор задачи (UUID).
* `status`: Состояние задачи (pending, running, completed, deleted).
* `created_at`: Время создания задачи.
* `duration`: Продолжительность выполнения

## Примеры запросов с curl

1. Создать задачу
```bash
curl -X POST http://localhost:8080/task
```

#### Пример ответа:
```json
{
  "id":"c5c23e5c-3ef4-4ee5-aa68-22fb0f2f566f"
}
```

2. Получить статус задачи
```bash
curl "http://localhost:8080/task/?id=c5c23e5c-3ef4-4ee5-aa68-22fb0f2f566f"
```

#### Пример ответа (для running):
```json
{
  "id": "c5c23e5c-3ef4-4ee5-aa68-22fb0f2f566f",
  "status": "running",
  "created_at": "2025-06-29T22:23:33+03:00",
  "duration": "3m20s"
}
```

#### Пример ответа (для completed):
```json
{
  "id": "c5c23e5c-3ef4-4ee5-aa68-22fb0f2f566f",
  "status": "completed",
  "created_at": "2025-06-29T22:23:33+03:00",
  "duration": "3m20s"
}
```

3. Удалить задачу
```bash
curl -X DELETE "http://localhost:8080/task/?id=450cf71a-f204-4eec-8f87-2453380ca0a5"
```

#### Пример ответа:
```json
{
  "message": "task deleted"
}
```

### Если задачи нет:
```json
{
  "Task not found"
}
```
