# High-Performance RESTful API with Golang 🐹

[![Go Report Card](https://goreportcard.com/badge/github.com/abedehakims/Rest-API-Golang)](https://goreportcard.com/report/github.com/abedehakims/Rest-API-Golang)
![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat&logo=go)
![Architecture](https://img.shields.io/badge/Architecture-Clean%20Architecture-green)

## 📖 Introduction
This repository contains a robust and scalable RESTful API built using **Go (Golang)**. The project is designed with a focus on high performance, maintainability, and clean code principles. It serves as a backend foundation for handling complex data transactions with minimal latency.

## 🛠 Features
* **RESTful Endpoints:** Full implementation of standard HTTP methods (GET, POST, PUT, DELETE).
* **Middleware Integration:** Includes logging, recovery, and CORS handling.
* **Authentication:** [Optional: JWT (JSON Web Token) based authentication for secure access].
* **Database Integration:** Efficient data persistence using GORM MySQL.
* **JSON Serialization:** High-speed request/response handling.
* **Input Validation:** Strict data validation to ensure API reliability.

## 🏗 Project Structure
This project follows a modular structure (inspired by Clean Architecture) to separate concerns:
```text
.
├── config/         # Database and environment configurations
├── controllers/    # Request handlers and logic orchestration
├── models/         # Database schemas and data structures
├── routes/         # API route definitions
├── middleware/     # Custom middleware (Auth, Logger)
└── main.go         # Application entry point

## 🚀 Technical Stack
```text
Language: Go (Golang)
Web Framework: Echo
Database: MySQL
Documentation: Youtube and Echo framework Docs

## 📡 API Endpoints
Method,Endpoint,Description,Auth Required
GET,/api/v1/items,Fetch all records,No
POST,/api/v1/items,Create a new record,Yes
GET,/api/v1/items/:id,Get details of a specific record,No
PUT,/api/v1/items/:id,Update an existing record,Yes
DELETE,/api/v1/items/:id,Remove a record,Yes

## 🏁 Getting Started
  * **Prerequisites**
    1. Go installed (version 1.20 or later)
    2. A running instance of (Database Name)

  * **Installation**
  ```bash
  git clone [https://github.com/abedehakims/Rest-API-Golang.git](https://github.com/abedehakims/Rest-API-Golang.git)

  * **Install dependencies:**
  ```bash
  go mod tidy
  Set up Environment Variables: Create a .env file and configure your database URL and server port.

  * **Run the server:**
  ```bash
  go run main.go

  * **🧪 Testing**
  To run the automated tests for this API:

  ```bash
  go test ./... -v

**Developed with ❤️ by Abede Hakims**
