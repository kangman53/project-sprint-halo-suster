package file_service

import (
	"github.com/gofiber/fiber/v2"
)

type FileService interface {
	Upload(ctx *fiber.Ctx) (string, error)
}
