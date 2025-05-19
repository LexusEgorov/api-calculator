# api-calculator

Проект 1: Калькулятор с REST API

## Обязательные заголовки
Authorization: <uID> - разделение запросов от пользователей по идентификатору

## Описание ручек

POST (application/json)

localhost:8080/sum
localhost:8080/mult

Принимают числа в формате a, b, c...
Возвращают результат сложения/умножения для данных чисел

GET
localhost:8080/history

Возвращает историю запросов для <uID>

## Примеры запросов и ответы
Сложение:

curl --location 'localhost:8080/sum' \
--header 'Authorization: 111' \
--header 'Content-Type: application/json' \
--data '{
    "input": "1, 2"
}'

{
    "input": "1, 2",
    "action": "SUM",
    "result": 3
}

Умножение:

curl --location 'localhost:8080/mult' \
--header 'Authorization: 111' \
--header 'Content-Type: application/json' \
--data '{
    "input": "1, 2"
}'

{
    "input": "1, 2",
    "action": "MULT",
    "result": 2
}

История:

curl --location 'localhost:8080/history' \
--header 'Authorization: 111'

[
    {
        "input": "1, 2",
        "action": "MULT",
        "result": 2
    },
    {
        "input": "1, 2",
        "action": "SUM",
        "result": 3
    }
]