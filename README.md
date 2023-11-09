# ORDER SERVICE

## Общая информация:
Реализация сервиса для сохранения orders(заказов) из NATS-streaming в PostgreSQL. 

# Как поднять:

## 1. Запуск контейнеров с nats-streaming и PostgreSQL:
```
make compose
```
## 2. Создание таблицы в поднятой в Docker БД:
```
make upmigrate
```
## 3. Запуск сервера:
```
go run cmd/main.go
```
## 4. Автоматическая публикация данных:
```
go run cmd/publisher/publisher.go
```
## 5. Остановка контейнеров:
```
make stop
```

## Стек технологий проекта:
* GO
* NATS-streaming
* Postgres
* Docker
* Goqu
* Gin
* HTML
* CSS
* Javascript
