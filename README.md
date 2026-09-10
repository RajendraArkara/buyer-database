# Buyer Database API

A backend REST API for managing buyer (customer) data and user authentication, built in Go. Originally developed to support buyer/customer data management for **CV Arkara Surya Abadi**, an export business for rattan and woven handicrafts.

## Features
- User authentication (sign up, login, forgot password) secured with JWT
- Buyer CRUD (create, read, update, delete) endpoints
- JWT-based auth middleware protecting sensitive routes
- Clean architecture: handler → usecase → repository layers

## Tech Stack
- **Go**
- **Gin** — HTTP web framework
- **PostgreSQL** (via `pgx` driver)
- **JWT** (`golang-jwt/jwt`) for authentication

## API Endpoints

### User
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/user/create-user` | Register a new user |
| POST | `/user/login` | Log in and receive a JWT token |
| PATCH | `/user/forgot-password` | Reset password |
| GET | `/user/get-all` | List all users |

### Buyer
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/buyer` | List all buyers |
| GET | `/buyer/:id` | Get a buyer by ID |
| POST | `/buyer/create-buyer` | Create a new buyer *(requires JWT)* |
| PATCH | `/buyer/update/:id` | Update a buyer |
| DELETE | `/buyer/delete/:id` | Delete a buyer |

## Architecture

The project follows a clean architecture pattern, separating concerns into:

- `handler/` — HTTP layer: routes and request/response handling
- `internal/usecase/` — business logic layer
- `infrastructure/repository/` — data access layer (PostgreSQL)
- `middlewares/` — JWT authentication middleware

## Running Locally

1. Clone this repository
2. Create a PostgreSQL database named `buyer-database`
3. Update the connection string in `infrastructure/db/db.go` to match your local setup
4. Install dependencies:
   ```bash
   go mod tidy
   ```
5. Run the server:
   ```bash
   go run main.go
   ```

The API will be available at `localhost:8080`.

## Background

This project started as a way to manage buyer/customer records for a real export business, and as hands-on practice applying clean architecture and JWT authentication in Go.
