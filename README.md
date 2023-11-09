# ORDER SERVICE

# Запуск контейнеров с nats-streaming и PostgreSQL:
```
docker-compose up -d
```
# Создание таблицы в поднятой в Docker БД:
```
migrate -path internal/db/migration -database 'postgres://wb:password@localhost:5040/orderdb?sslmode=disable' up
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
docker stop orders-service && docker stop nats-stream
```
