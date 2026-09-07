# learn-vibe-code (PPOB Backend Service)

Backend service for PPOB (Payment Point Online Bank) built with **Go**, **Gin Web Framework**, and **GORM (PostgreSQL)**.

---

## 🛠 Tech Stack
- **Language**: Go 1.26+
- **Web Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database Driver**: [GORM PostgreSQL Driver](https://github.com/go-gorm/postgres)
- **Security / Encryption**: `golang.org/x/crypto/bcrypt`
- **ID Generation**: `github.com/google/uuid` (UUID v4)
- **Configuration**: [godotenv](https://github.com/joho/godotenv)
- **Logging**: Standard Library `log/slog` (Structured JSON logging)
- **Deployment**: [Render.com](https://render.com) (`render.yaml` Blueprint & Dockerfile)

---

## 📁 Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go          # Application entry point & graceful shutdown
├── internal/
│   ├── config/
│   │   └── config.go        # Environment variable & Render DATABASE_URL loader
│   ├── database/
│   │   └── postgres.go      # GORM PostgreSQL connection & AutoMigrate
│   ├── dto/
│   │   ├── response.go      # Standardized API response format
│   │   └── user_dto.go      # Request DTOs & validations
│   ├── handler/
│   │   ├── health.go        # Health check handler
│   │   └── user_handler.go  # User registration handler
│   ├── model/
│   │   └── user.go          # User entity model (UUID PK)
│   ├── repository/
│   │   └── user_repository.go # Database queries
│   ├── service/
│   │   └── user_service.go  # Business logic & password hashing
│   └── server/
│       └── router.go        # Gin router & route registration
├── Dockerfile               # Multi-stage container build
├── .dockerignore
├── render.yaml              # Render.com Infrastructure as Code blueprint
├── .env.example             # Example environment variables
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module definitions
├── go.sum                   # Go dependencies checksums
└── README.md                # Project documentation
```

---

## 🚀 Getting Started Locally

### 1. Prerequisites
- **Go** (version 1.26 or higher recommended)
- **PostgreSQL** database server

### 2. Environment Configuration
Copy the `.env.example` file to `.env` and adjust the variables to your local database configuration:
```bash
cp .env.example .env
```

Default configuration in `.env.example`:
```ini
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ppob_db
DB_SSLMODE=disable
```

### 3. Install Dependencies
```bash
go mod download
```

### 4. Run the Server
```bash
go run ./cmd/api
```

---

## ☁️ Deployment to Render.com

This repository includes a `render.yaml` blueprint file for zero-config deployment on Render.

### Option 1: Blueprint Deployment (Recommended)
1. Push your repository to GitHub.
2. Log in to [Render Dashboard](https://dashboard.render.com/).
3. Click **New +** -> **Blueprint**.
4. Connect this repository (`learn-vibe-code`).
5. Render will automatically provision:
   - PostgreSQL Database (`ppob-postgres`)
   - Go Web Service (`learn-vibe-code-api`) with automatic `DATABASE_URL` binding.
6. Click **Apply**.

### Option 2: Manual Web Service
- **Runtime**: Go or Docker
- **Build Command**: `go build -o bin/api ./cmd/api`
- **Start Command**: `./bin/api`
- **Environment Variables**:
  - `APP_ENV`: `production`
  - `DATABASE_URL`: (Connection string from your Render PostgreSQL instance)

---

## 📡 API Endpoints

### Standard Response Format
All JSON responses follow this unified format:
```json
{
  "message": "Declarative description",
  "result": "ok | error",
  "data": null
}
```

### Endpoint List

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/health` | Health Check (Server & Database Status) |
| `GET` | `/api/health` | API Health Check |
| `POST` | `/api/register` | User Registration |

#### Register User Example (`POST /api/register`):
**Request:**
```json
{
  "name": "darsam",
  "email": "darsam@gmail.com",
  "password": "password123"
}
```

**Success Response (201 Created):**
```json
{
  "message": "User registered successfully",
  "result": "ok",
  "data": null
}
```
