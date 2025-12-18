package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"user-api-go/internal/repository"
	"user-api-go/internal/service"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var body struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	dob, err := time.Parse("2006-01-02", body.Dob)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "dob must be YYYY-MM-DD"})
	}

	user, err := h.repo.Create(c.Context(), body.Name, dob)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob,
		"age":  service.CalculateAge(user.Dob),
	})
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.repo.Get(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob,
		"age":  service.CalculateAge(user.Dob),
	})
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.repo.List(c.Context())
	if err != nil {
		return c.Status(500).JSON(err)
	}

	resp := make([]fiber.Map, 0)
	for _, u := range users {
		resp = append(resp, fiber.Map{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob,
			"age":  service.CalculateAge(u.Dob),
		})
	}

	return c.JSON(resp)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var body struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	dob, err := time.Parse("2006-01-02", body.Dob)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "dob must be YYYY-MM-DD"})
	}

	user, err := h.repo.Update(c.Context(), int32(id), body.Name, dob)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob,
		"age":  service.CalculateAge(user.Dob),
	})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.repo.Delete(c.Context(), int32(id)); err != nil {
		return c.Status(500).JSON(err)
	}

	return c.SendStatus(204)
}
