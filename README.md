# TestTaskBrandScout
Test Golang task for BrandScout

## Тестовое задание:
REST API-сервис для работы с цитатами. Реализованы следующие функции:

- Получение всех цитат
- Фильтрация цитат по автору
- Получение случайной цитаты
- Создание новой цитаты
- Удаление цитаты по ID

Unit-тесты:
- go test ./internal/handler -v - Тестирование хендлеров
- go test ./internal/service -v - Тестирование сервисного слоя
*Перед unit-тестами необходимо запустить контейнер database с базой данных, поменять database на localhost в .env и выполнить goose up*
## Быстрый старт


```bash
git clone https://github.com/kiriksik/TestTaskBrandScout.git
cd TestTaskBrandScout
docker-compose up --build
```

Список команд, которые проверялись через curl:
```bash 
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Kokorin", "quote":"A chto pohozhe."}'
curl http://localhost:8080/quotes 
curl http://localhost:8080/quotes/random
curl http://localhost:8080/quotes?author=Confucius
curl -X DELETE http://localhost:8080/quotes/1
```

## Технологии
По заданию используются стандартные библиотеки(требование), кроме:
- github.com/lib/pq - драйвер для БД postgreSQL
- github.com/joho/godotenv - для удобной подгрузки .env из корня

В дополнение к заданию мной было реализованно:
- Хранение цитат в базе данных PostgreSQL
- - Для миграций использовался Goose
- - Для автогенерации запросов SQLC
- Добавлен Dockerfile и docker-compose для облегчения сборки вместе с БД
