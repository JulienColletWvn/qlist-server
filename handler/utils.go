package handler

import (
	"github.com/gofiber/fiber"
)

type Handler func(c *fiber.Ctx) error
