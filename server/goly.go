package server

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tajpouria/goly/repository"
	"gorm.io/gorm"
)

func getGoliesHandler(ctx *fiber.Ctx) error {
	golies, err := repository.Model.GetGolies()
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.ErrInternalServerError.Code).
			JSON(MsgInternalServerErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(golies)
}

func getGolyHandler(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(MsgInternalServerErr)
	}

	goly, err := repository.Model.GetGoly(id)
	if err != nil {
		if err != nil {
			switch err.Error() {
			case gorm.ErrRecordNotFound.Error():
				return ctx.Status(fiber.StatusNotFound).
					JSON(fiber.Map{"message": fmt.Sprintf("There is no Goly with id: %v", id)})
			default:
				log.Error(err)
				return ctx.Status(fiber.StatusInternalServerError).
					JSON(MsgInternalServerErr)
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func RegisterGolyRoute(router fiber.Router) {
	router.Get("", getGoliesHandler)
	router.Get("/:id", getGolyHandler)
}
