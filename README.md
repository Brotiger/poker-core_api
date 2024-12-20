# Расписной покер (Core API)
## Зависомости
* go1.22
* make

## Документация
* docs
    * mongodb
        * collection - _структура коллекций_
            * users.js


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
```