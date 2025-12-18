package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLogger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		requestID := uuid.New().String()
		c.Set("X-Request-ID", requestID)

		err := c.Next()

		log.Info("request completed",
			zap.String("request_id", requestID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("duration", time.Since(start)),
			zap.Int("status", c.Response().StatusCode()),
		)

		return err
	}
}
