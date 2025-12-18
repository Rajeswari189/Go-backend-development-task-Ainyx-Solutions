# Go-backend-development-task-Ainyx-Solutions

# User API – Go Backend Assessment

## Tech Stack
- Go (Golang)
- GoFiber
- PostgreSQL
- SQLC
- Uber Zap

## Features
- Create user with name and dob
- Fetch user with dynamically calculated age
- Clean architecture (handler → service → repository)

## Setup

1. Start PostgreSQL
2. Run migrations
3. Generate SQLC code
   sqlc generate
4. Run server
   go run cmd/server/main.go

Server runs on http://localhost:8080
