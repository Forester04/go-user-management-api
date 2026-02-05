# User Management API (Go + GORM)
A robust, production-ready RESTful API for User Management built with Go. This project demonstrates Clean Architecture principles, featuring JWT authentication, soft deletes, and a centralized error-handling system.

## 🚀 Key Features
Authentication & Security: Secure Register/Login flow using JWT (JSON Web Tokens) and Bcrypt for password hashing.

## Full CRUD Operations:

 - Profile retrieval and updates.

 - Soft Deletes: Users are flagged as deleted in the database without losing record integrity.

 - Architecture & Design:

 - DTOs (Data Transfer Objects): Strict input validation.

 - ViewModels: Controlled response structures to prevent sensitive data leakage.

 - Repository Pattern: Decoupled business logic from database implementation (GORM).

 - Advanced Error Handling: Implementation of GoCleanError for consistent, centralized, and clear API error responses.

 - Middleware: Custom recovery, logging, and JWT-based authorization gates.

## 🛠 Tech Stack
 - Language: Go (Golang)

 - Framework: Gin Gonic

 - ORM: GORM

 - Database: PostgreSQL / MySQL

 - Auth: JWT-Go
### 🚦 Getting Started
 - Prerequisites: Go 1.21+

## 📂 Project Structure

```text
├── cmd/                # Application entry point
├── internal/
│   ├── delivery/       # Handlers/Controllers & Middleware
│   ├── domain/         # Interfaces and Models
│   ├── repository/     # GORM implementation
│   ├── service/        # Business logic
│   └── dto/            # Input/Output Structs (DTOs & ViewModels)
├── pkg/
│   └── gocleanerror/   # Centralized error handling utility
└── main.go
