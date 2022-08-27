package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tajpouria/goly/repository"
	"github.com/tajpouria/goly/utils"
	"gorm.io/gorm"
)

func getGoliesHandler(ctx *fiber.Ctx) error {
	golies, err := repository.Model.GetGolies()
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(msgInternalServerErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(golies)
}

func getGolyHandler(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
	}

	goly, err := repository.Model.GetGoly(id)
	if err != nil {
		if err != nil {
			switch err.Error() {
			case gorm.ErrRecordNotFound.Error():
				return ctx.Status(fiber.StatusNotFound).JSON(msgGolyIDNotFound(id))
			default:
				log.Error(err)
				return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func createGolyHandler(ctx *fiber.Ctx) error {
	goly := new(repository.Goly)

	if err := ctx.BodyParser(goly); err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
	}

	if err := repository.Validate.Struct(goly); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrResponse(err))
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(16)
	}

	if err := repository.Model.CreateGoly(&repository.Goly{
		Redirect: goly.Redirect,
		Goly:     goly.Goly,
		Random:   goly.Random,
		Clicked:  0,
	}); err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
	}

	return ctx.Status(fiber.StatusCreated).JSON(goly)
}

func updateGolyHandler(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
	}
	goly := new(repository.Goly)

	if err := ctx.BodyParser(goly); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.ErrResponse(err))
	}

	if err := repository.Model.UpdateGoly(&repository.Goly{
		ID:       id,
		Redirect: goly.Redirect,
		Goly:     goly.Goly,
		Random:   goly.Random,
	}); err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(goly)
}

func deleteGolyHandler(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
	}

	if err := repository.Model.DeleteGoly(&repository.Goly{ID: id}); err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func RegisterGolyRoute(router fiber.Router) {
	router.Get("", getGoliesHandler)
	router.Get("/:id", getGolyHandler)
	router.Post("", createGolyHandler)
	router.Put("/:id", updateGolyHandler)
	router.Delete("/:id", deleteGolyHandler)
}
