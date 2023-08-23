# Test task BackDev

## Описание
Проeк "Test task BackDev" представляет собой часть сервиса аутентификации, в котором представлены два REST маршрута. Первый маршрут выдает пару Access, Refresh токенов для пользователя с идентификатором указанным в параметре запроса. Второй маршрут выполняет Refresh операцию на пару Access, Refresh токенов.

## Как использовать
1. Запуск сервера:
    - Установите Go, если его еще нет: [Установка Go](https://golang.org/doc/install)
    - Клонируйте репозиторий проекта:
        ```
        git clone https://github.com/bilorukavsky/Test_task_BackDev
        ```
    - Перейдите в директорию проекта:
        ```
        cd Test_task_BackDev
        ```
    - Запустите сервер:
        ```
        go run main.go
        ```
    - Сервер будет запущен и будет слушать на порту 8080.
2. Использование API:  
    - Чтобы получить пару токенов, отправьте POST-запрос на `/login`
    - Тело запроса должно быть в формате JSON и содержать поле `username` со значением формата `string`.
    - Пример запроса с использованием `curl`:
        ```
        curl -X POST -H "Content-Type: application/json" -d '{"url": "1"}' http://localhost:8080/login
        ```
    - Ответ будет содержать два токена в полях `access_token` и `refresh_token`:
        ```json
        {
          "access_token":"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI3ODAzNDgsInN1YiI6IjEifQ.TQow1HO7ts1dA4jm4OabNPXT8v5Ex7ERezwgtTB3FvjcToT9UJ02_XdfdiS1kRKm-1tb6eIJS3mQ1ioinKgGqg",
          "refresh_token":"yMNNk1niPswfq-CwaotE8H6L1cRWTBWDy1hdXG9GKjA"
        }
        ```
  
    Методы API
    ---
    ### POST /login
    Выдает access и refresh токены.

    Запрос:
    - Тело запроса должно быть в формате JSON.
    - Поле `username` должно содержать строку.

    Пример запроса:
    ```json
    {
    "username": "1"
    }
    ```
    Ответ:
    - Ответ будет в формате JSON.
    - Поля `access_token` и `refresh_token` будет содержать токены.

    Пример ответа:
    ```json
    {
        "access_token":"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI3ODAzNDgsInN1YiI6IjEifQ.TQow1HO7ts1dA4jm4OabNPXT8v5Ex7ERezwgtTB3FvjcToT9UJ02_XdfdiS1kRKm-1tb6eIJS3mQ1ioinKgGqg",
        "refresh_token":"yMNNk1niPswfq-CwaotE8H6L1cRWTBWDy1hdXG9GKjA"
    }
    ```
    ### POST /refresh
    Выдает обновленные access и refresh токены.

    Запрос:
    - Тело запроса должно быть в формате JSON.
    - Поле `username` должно содержать строку, поле `refresh_token` содержить токен 
    Пример запроса:
    ```json
    {
    "username": "1",
    "refresh_token":"yMNNk1niPswfq-CwaotE8H6L1cRWTBWDy1hdXG9GKjA"
    }
    ```
    Ответ:
    - Ответ будет в формате JSON.
    - Поля `access_token` и `refresh_token` будет содержать токены.

    Пример ответа:
    ```json
    {
        "access_token":"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI3ODAzNDgsInN1YiI6IjEifQ.TQow1HO7ts1dA4jm4OabNPXT8v5Ex7ERezwgtTB3FvjcToT9UJ02_XdfdiS1kRKm-1tf62IJS3mQ1ioidKtGqg",
        "refresh_token":"yMNNk1niPswfq-CwaotE8H6L1cRWTBWDy1hdXG9GKjA"
    }
    ```
## Доплнительная информация
Проект "Test task backDev" - это простой пример реализации части сервиса аутентификации с использованием языка программирования Go и веб-сервера. Он может служить основой для разработки более сложных и функциональных систем аунтификации.
    