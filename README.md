# ORDER SERVICE

# Запуск контейнеров с nats-streaming и PostgreSQL:
```
make compose
```
# Создание таблицы в поднятой в Docker БД:
```
make upmigrate
```
# Запуск сервера:
```
go run cmd/main.go
```
# Автоматическая публикация данных:
```
go run cmd/publisher/publisher.go
```
# Остановка контейнеров:
```
make stop
```
