# learn-vibe-code (PPOB Backend Service)

Backend service for PPOB (Payment Point Online Bank) built with **Go**, **Gin Web Framework**, and **GORM (PostgreSQL)**.

---

## 🛠 Tech Stack
- **Language**: Go 1.25+
- **Web Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database Driver**: [GORM PostgreSQL Driver](https://github.com/go-gorm/postgres)
- **Configuration**: [godotenv](https://github.com/joho/godotenv)
- **Logging**: Standard Library `log/slog` (Structured JSON logging)

---

## 📁 Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go          # Application entry point & graceful shutdown
├── internal/
│   ├── config/
│   │   └── config.go        # Environment variable loader
│   ├── database/
│   │   └── postgres.go      # GORM PostgreSQL connection setup
│   ├── handler/
│   │   └── health.go        # HTTP request handlers (e.g. Health check)
│   └── server/
│       └── router.go        # Gin router & middleware configuration
├── .env.example             # Example environment variables
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module definitions
├── go.sum                   # Go dependencies checksums
└── README.md                # Project documentation
```

---

## 🚀 Getting Started

### 1. Prerequisites
- **Go** (version 1.25 or higher recommended)
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

## 📡 Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/health` | Application & Database Health Check |
| `GET` | `/api/v1/health` | API v1 Health Check |

**Sample Response (`/health`):**
```json
{
  "database": "connected",
  "status": "ok",
  "timestamp": "2026-09-06T12:00:00Z"
}
```
