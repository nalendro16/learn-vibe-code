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

## ☁️ Deployment

### 1. Deploy to Vercel (Serverless Go)
Proyek ini sudah dilengkapi dengan `vercel.json` dan `api/index.go` sehingga dapat langsung di-deploy ke **Vercel**:

1. Pastikan Anda memiliki database PostgreSQL cloud (misalnya dari **[Neon.tech](https://neon.tech)**, **[Supabase](https://supabase.com)**, atau **Vercel Postgres**).
2. Install Vercel CLI atau deploy via Dashboard:
   - **Via Vercel Dashboard**:
     - Import repository GitHub ini ke Vercel.
     - Tambahkan **Environment Variables** di project settings:
       - `DATABASE_URL`: `postgres://user:password@ep-xyz.neon.tech/neondb?sslmode=require`
       - `APP_ENV`: `production`
     - Klik **Deploy**.
   - **Via CLI**:
     ```bash
     npm i -g vercel
     vercel
     ```

---

### 2. Deploy to Render.com
Repository ini juga menyertakan file blueprint `render.yaml` untuk deploy di Render:

- **Via Blueprint**:
  1. Buka [Render Dashboard](https://dashboard.render.com/) -> **New +** -> **Blueprint**.
  2. Pilih repository ini dan klik **Apply**. Render otomatis membuat database dan web service.
- **Via Manual Web Service**:
  - Build Command: `go build -o bin/api ./cmd/api`
  - Start Command: `./bin/api`
  - Environment Variables: `APP_ENV=production`, `DATABASE_URL=...`

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
