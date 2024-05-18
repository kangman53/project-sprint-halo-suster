package controller

import (
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	file_service "github.com/kangman53/project-sprint-halo-suster/service/file"

	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	FileService file_service.FileService
}

func NewFileController(fileService file_service.FileService) *FileController {
	return &FileController{
		FileService: fileService,
	}
}

func (controller FileController) Upload(ctx *fiber.Ctx) error {
	resp, err := controller.FileService.Upload(ctx)
	if err != nil {
		return exc.Exception(ctx, err)
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
