package file_service

import (
	"github.com/gofiber/fiber/v2"
	file_entity "github.com/kangman53/project-sprint-halo-suster/entity/file"
)

type FileService interface {
	Upload(ctx *fiber.Ctx) (file_entity.UploadImageResponse, error)
}
