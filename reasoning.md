# Reasoning & Design Decisions

## Architecture
I used a layered architecture:
- Handler: HTTP request/response
- Service: business logic (age calculation)
- Repository: database access using SQLC

Each layer has a single responsibility, improving readability and maintainability.

## Why SQLC
SQLC provides type-safe database access and avoids runtime SQL errors while keeping full SQL control.

## Age Calculation
Age is not stored in the database.
It is calculated dynamically in Go using the time package to ensure accuracy.

## Error Handling
- 400 for invalid input
- 404 when user not found
- 500 for internal errors

## Conclusion
This design is simple, clean, testable, and follows Go backend best practices as of my best knowledge 








