package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"user-api-go/config"
	"user-api-go/db/sqlc"
	"user-api-go/internal/handler"
	"user-api-go/internal/logger"
	"user-api-go/internal/middleware"
	"user-api-go/internal/repository"
	"user-api-go/internal/routes"
)

func main() {
	cfg := config.Load()
	logg := logger.New()

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	q := sqlc.New(db)
	repo := repository.NewUserRepository(q)
	h := handler.NewUserHandler(repo)

	app := fiber.New()
	app.Use(middleware.RequestLogger(logg))

	routes.Register(app, h)

	log.Fatal(app.Listen(":" + cfg.Port))
}
