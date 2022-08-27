package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var msgInternalServerErr = fiber.Map{"message": "oops! something went wrong."}

func msgGolyIDNotFound(id uint64) fiber.Map {
	return fiber.Map{"message": fmt.Sprintf("There is no any Goly with id: '%v'.", id)}
}

func msgGolyURLNotFound(url string) fiber.Map {
	return fiber.Map{"message": fmt.Sprintf("There is no any Goly with URL: '%s'.", url)}
}
