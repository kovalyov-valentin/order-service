# ORDER SERVICE

## Общая информация:
Реализация сервиса для сохранения orders(заказов) из NATS-streaming в PostgreSQL. 

# Как поднять:
<ol>
<li>## Запуск контейнеров с nats-streaming и PostgreSQL:
```
make compose
```
<li>## Создание таблицы в поднятой в Docker БД:
```
make upmigrate
```
<li>## Запуск сервера:
```
go run cmd/main.go
```
<li>## Автоматическая публикация данных:
```
go run cmd/publisher/publisher.go
```
<li>## Остановка контейнеров:
```
make stop
```
</ol>
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
