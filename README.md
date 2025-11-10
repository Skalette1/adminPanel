# Admin Panel API

**Admin Panel** — это бэкенд для управления пользователями и ролями, реализованный на **Go** с использованием **Gin** и **PostgreSQL**.

API документировано через **Swagger / OpenAPI**.

---

## 📂 Структура проекта

```
adminPanel/
├── main.go                    # Точка входа
├── go.mod
├── internal/
│   ├── db/                    # Подключение к БД
│   ├── handlers/              # API handlers
│   ├── models/                # Сущности
│   ├── repository/            # Репозитории (CRUD)
│   └── routes/                # Роутеры
├── dto/                       # DTO для запросов и ответов
├── openapi/                   # OpenAPI / Swagger YAML
└── docs/                      # Swagger UI
```

---

## ⚡ Установка и запуск

### 1. Клонирование

```bash
git clone https://github.com/Skalette1/adminPanel.git
cd adminPanel
```

### 2. Установка зависимостей

```bash
go mod download
```

### 3. Настройка базы данных

Создайте PostgreSQL и укажите параметры в `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=adminpanel
```

Или используйте Docker:

```bash
docker run --name adminpanel-db -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=adminpanel -p 5432:5432 -d postgres
```

### 4. Запуск приложения

**Локально:**

```bash
go run main.go
```

**Через Docker:**

```bash
docker build -t adminpanel .
docker run -p 8080:8080 adminpanel
```

**Через Docker Compose:**

```yaml
version: "3.9"
services:
  db:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: adminpanel
    ports:
      - "5432:5432"
  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - db
```

```bash
docker compose up
```

---

## 📖 API

Документация генерируется через **Swagger**.
Доступ к Swagger UI:

```
http://localhost:8080/swagger/index.html
```

---

### 👤 Users

| Метод  | URL         | Описание                    | Статусы            |
| ------ | ----------- | --------------------------- | ------------------ |
| POST   | /users      | Создать пользователя        | 201, 400, 500      |
| GET    | /users      | Получить всех пользователей | 200, 500           |
| GET    | /users/{id} | Получить пользователя по ID | 200, 404, 500      |
| PUT    | /users/{id} | Обновить пользователя       | 200, 400, 404, 500 |
| DELETE | /users/{id} | Удалить пользователя        | 200, 404, 500      |

Пример запроса (POST `/users`):

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "ivan",
    "email": "ivan@example.com",
    "password": "secret",
    "role_id": 1
}'
```

Пример успешного ответа:

```json
{
  "id": 1,
  "username": "ivan",
  "email": "ivan@example.com",
  "role_id": 1,
  "is_active": true,
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T01:00:00Z"
}
```

Пример ошибки:

```json
{
  "message": "Invalid input",
  "details": "Email is required"
}
```

---

### 🛡 Roles

| Метод  | URL         | Описание            | Статусы            |
| ------ | ----------- | ------------------- | ------------------ |
| POST   | /roles      | Создать роль        | 201, 400, 500      |
| GET    | /roles      | Получить все роли   | 200, 500           |
| GET    | /roles/{id} | Получить роль по ID | 200, 404, 500      |
| PUT    | /roles/{id} | Обновить роль       | 200, 400, 404, 500 |
| DELETE | /roles/{id} | Удалить роль        | 200, 404, 500      |

Пример запроса (POST `/roles`):

```bash
curl -X POST http://localhost:8080/roles \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "permission": "full_access"
}'
```

Пример успешного ответа:

```json
{
  "id": 1,
  "username": "admin",
  "permission": "full_access",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T01:00:00Z"
}
```

---

### 📦 DTO (пример)

```go
type CreateUserRequest struct {
    Username string `json:"username" binding:"required" example:"ivan"`
    Email    string `json:"email" binding:"required,email" example:"ivan@example.com"`
    Password string `json:"password" binding:"required" example:"secret"`
    RoleId   int    `json:"role_id,omitempty" example:"1"`
}

type UserSuccessResponse struct {
    ID        int       `json:"id" example:"1"`
    Username  string    `json:"username" example:"ivan"`
    Email     string    `json:"email" example:"ivan@example.com"`
    RoleId    int       `json:"role_id" example:"1"`
    IsActive  bool      `json:"is_active" example:"true"`
    CreatedAt time.Time `json:"created_at" example:"2025-01-01T00:00:00Z"`
    UpdatedAt time.Time `json:"updated_at" example:"2025-01-01T01:00:00Z"`
}

type UserErrorResponse struct {
    Message string `json:"message" example:"Invalid input"`
    Details string `json:"details" example:"Email is required"`
}
```

