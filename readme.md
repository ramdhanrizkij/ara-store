# Ara Store — Backend API (Single Vendor E-Commerce)

Ara Store is a backend API project built with Go, following Clean Architecture and Domain-Driven Design principles. This project is designed as a learning platform to build a scalable, maintainable, and production-ready e-commerce system.

---

## 🚀 Tech Stack

* **Language**: Go 1.25+
* **Framework**: Gin
* **ORM**: GORM v2
* **Database**: PostgreSQL 16
* **Migration**: Goose
* **Queue**: Asynq (Redis)
* **Cron Jobs**: robfig/cron
* **Storage**: AWS S3
* **Auth**: JWT
* **Validation**: go-playground/validator
* **Logger**: Zap
* **Config**: Viper + .env

---

## 📁 Project Structure

```
cmd/
  server/         # API entry point
  worker/         # Background worker

internal/
  config/         # Configuration loader
  logger/         # Logger setup
  database/       # DB connection & migration

  domain/         # Entities & interfaces
  application/    # Business logic (services)
  infrastructure/ # External systems (DB, S3, queue, cron)
  interfaces/     # HTTP layer (handlers, middleware, router)

pkg/
  apperror/       # Custom error handling
  validator/      # Validation wrapper
  utils/          # Helper functions

migrations/       # SQL migration files
```

---

## ⚙️ Setup & Installation

### 1. Clone Project

```bash
git clone https://github.com/ramdhanrizkij/ara-store.git
cd ara-store
```

---

### 2. Setup Environment

Copy env file:

```bash
cp .env.example .env
```

---

### 3. Install Dependencies

```bash
go mod tidy
```

---

### 4. Run Docker Services

```bash
make docker-up
```

Services:

* PostgreSQL → `localhost:5432`
* Redis → `localhost:6379`

---

### 5. Run Migrations

```bash
make migrate-up
```

---

### 6. Run Server

```bash
make run
```

Server will run at:

```
http://localhost:8080
```

---

### 7. Run Worker (Async Jobs)

```bash
make worker
```

---

## 🧪 Testing

```bash
make test
```

---

## 🛠️ Available Commands

| Command             | Description              |
| ------------------- | ------------------------ |
| `make run`          | Run API server           |
| `make build`        | Build binary             |
| `make worker`       | Run background worker    |
| `make test`         | Run tests                |
| `make migrate-up`   | Apply migrations         |
| `make migrate-down` | Rollback migrations      |
| `make docker-up`    | Start PostgreSQL & Redis |
| `make docker-down`  | Stop containers          |

---

## 🔐 Authentication

* JWT-based authentication
* Header format:

```
Authorization: Bearer <token>
```

Roles:

* `user`
* `admin`

---

## 📦 API Overview

Base URL:

```
/api/v1
```

### Auth

* POST `/auth/register`
* POST `/auth/login`
* GET `/auth/me`

### Products

* GET `/products`
* GET `/products/:id`
* POST `/products` (admin)

### Categories

* GET `/categories`
* POST `/categories` (admin)

### Orders

* POST `/orders`
* GET `/orders`

### Payments

* POST `/payments`

---

## 🧱 Architecture Principles

This project uses **Layered Clean Architecture**:

* **Handler** → HTTP layer (Gin)
* **Service** → Business logic
* **Repository** → Data access
* **Domain** → Core entities & contracts
* **Infrastructure** → External integrations

### Key Benefits

* Scalable
* Testable
* Maintainable
* Clear separation of concerns

---

## ⚡ Background Processing

### Queue (Asynq)

* Send emails
* Deduct stock
* Activity logging

### Cron Jobs

* Auto cancel unpaid orders
* Daily reports
* Cleanup soft deletes

---

## ☁️ File Storage (S3)

* Product images stored in AWS S3
* Supported formats: JPG, PNG, WEBP
* Max size: 5MB

---

## 📌 Development Notes

* Use UUID for all IDs
* Use soft delete (`deleted_at`)
* All responses follow standard wrapper:

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

---

## 🧠 Learning Goals

This project is designed to help you learn:

* Clean Architecture in Go
* REST API design
* Database schema design (PostgreSQL)
* Background jobs & async processing
* Production-grade backend practices

---

## 📄 License

This project is for educational purposes.
