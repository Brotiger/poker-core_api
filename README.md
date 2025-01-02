# Расписной покер (Core API)
Основное API игры

## Зависомости
* go1.22
* make

## Документация
* docs
    * mongodb
        * collection - _структура коллекций_
            * code.js
            * game.js
            * refresh_token.js
            * socket.js
            * user.js

## Запуск инфраструктуры
~~~bash
cd ./docker
docker compose up
~~~

## Сборка и запуск
### Конфигурация
_Прописать в `.env` параметры и экспортировать:_
```bash
cp .env.example .env
source .env
```

### Make команды
```bash
make build # сборка проекта
make up # запуск скомпилированного проекта
make run # запуск без компиляции

make dev # установка dev зависимостей
make docs # генерация swagger документации
make seed # генерация данных
```

## Ссылки
* [Клиентское приложение](https://github.com/Brotiger/poker-app)
* [Сервис отправки писем](https://github.com/Brotiger/poker-mailer)
* [WebSocket сервис](https://github.com/Brotiger/poker-websocket)