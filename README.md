# Quests
- [Сборка и запуск](#запуск)
- [Пути](#Пути)
## Сборка и запуск
1. Для сборки проекта нужно ввести команду в терминал находясь в папке Quest
```sh
docker compose build
```
2. Для запуска проекта нужно ввести команду в терминал находясь в папке Quest
```sh
docker compose up
```
3. Для запуска миграций нужно ввести команду
```sh
  docker run -v (path to folder)/quest/database/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:postgres@localhost:5432/quests?sslmode=disable" up
```
4. Команда для создания новых миграций
```sh
migrate create -ext sql -dir Database/Migrations/ -seq migration-name
```
## Пути

- [Add user](#add-user)
- [Users list](#users-list)
- [User history](#user-history)
- [Add quest](#add-quest)
- [Quest list](#quests-list)
- [Change quest](#change-quest)
- [Process signal](#process-signal)
- [Swagger](#swagger)

### Add User
- **Method:** POST
- **Path:** /user
- **Description:** Добавление нового пользователя.
- **Body example:**
```json
{
    "name" : "Some name",
    "balance": 1000
}
```

### Users list
- **Method:** GET
- **Path:** /user
- **Description:** Получение списка пользователей.

### User history
- **Method:** GET
- **Path:** /user/:id/history
- **Description:** Получение истории заданий пользователя.

### Add quest
- **Method:** POST
- **Path:** /quest
- **Description:** Добавление нового задания.
- **Body example:**
```json
{
    "name" : "Some name",
    "cost": 2000
}
```

### Quests list
- **Method:** GET
- **Path:** /quest
- **Description:** Получение списка заданий.

### Change quest
- **Method:** PUT
- **Path:** /quest/:id
- **Description:** Обновление задания.
- **Body example:**
```json
{
    "name" : "Some new name",
    "cost": 3000
}
```

### Process signal
- **Method:** POST
- **Path:** /signal
- **Description:** Выполнение задания пользователем.
- **Body example:**
```json
{
    "user_id" : 2,
    "quest_id": 3
}
```

### Swagger
- **Method:** get
- **Path:** /swagger/index.html
- **Description:**  Документация.


