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

## TestCases
<img width="832" height="154" alt="image" src="https://github.com/user-attachments/assets/2730a638-a633-4f36-b25e-89061432db6d" />

## Github push
git init
git status
git add .
git commit -m "Initial commit - User API with DOB and Age by Raji"
git remote add origin https://github.com/MYUSERNAME/user-api-go.git
git remote -v
git branch -M main
git push -u origin main

## To run App
go mod tidy
go run cmd/server/main.go

## To stop App
lsof -i :8080
Kill -9 <pid>
If crt+c not working

## CRUD Operations 

a. Create user ->

curl -X POST https://<codespace-name>-8080.app.github.dev/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","dob":"1998-05-21"}'

b. Get User ->

curl http://localhost:8080/users

c. Get User By ID ->

curl http://localhost:8080/users/1

d. Update User ->

curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","dob":"1997-04-10"}'

e. Delete User ->

curl -X DELETE http://localhost:8080/users/1



## We can use api tools like Postman for CRUD Operations 

## URL’s for route I used here

POST   /users
GET    /users
GET    /users/:id
PUT    /users/:id
DELETE /users/:id



