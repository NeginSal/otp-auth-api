# OTP Authentication API

A simple OTP-based authentication REST API built with Golang, Gin, MongoDB, JWT, and Docker.

## Features

- Send OTP to user phone number (simulated via console log)
- Verify OTP and generate JWT token
- Store users and OTP requests in MongoDB
- REST API documented with Swagger UI
- Dockerized for easy deployment with MongoDB

## Getting Started

### Prerequisites

- Go 1.21+
- Docker and Docker Compose (optional)
- MongoDB (if not using Docker)

### Running Locally (without Docker)

1. Set environment variables in `.env` file:

```
PORT=8080
MONGO_URI=mongodb://localhost:27017/otp_auth
JWT_SECRET=your_jwt_secret_key
OTP_EXPIRATION_MINUTES=5
```
2. Run the API server:  `go run cmd/main.go`


3. Access Swagger UI at:  `http://localhost:8080/swagger/index.html`

### Running with Docker
1. Build and start services:  `docker-compose up --build`
2. The API will be available at:  `http://localhost:8080`
3. Swagger UI at:  `http://localhost:8080/swagger/index.html`


### API Endpoints
- POST /send-otp - Send OTP to phone number
- POST /verify-otp - Verify OTP and receive JWT token

### Testing API
```
curl -X POST http://localhost:8080/send-otp \
  -H "Content-Type: application/json" \
  -d '{"phone":"09123456789"}'
```
