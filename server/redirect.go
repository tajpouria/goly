package server

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tajpouria/goly/repository"
	"gorm.io/gorm"
)

func redirectHandler(ctx *fiber.Ctx) error {
	url := ctx.Params("goly")
	goly, err := repository.Model.GetGolyByURL(url)
	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return ctx.Status(fiber.StatusNotFound).JSON(msgGolyURLNotFound(url))
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(msgInternalServerErr)
		}
	}

	goly.Clicked += 1
	if err := repository.Model.UpdateGoly(&goly); err != nil {
		log.Error(err)
	}

	return ctx.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func RegisterRedirectRouter(router fiber.Router) {
	router.Get("/:goly", redirectHandler)
}
