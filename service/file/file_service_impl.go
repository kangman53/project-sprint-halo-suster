package file_service

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	file_entity "github.com/kangman53/project-sprint-halo-suster/entity/file"
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	"github.com/kangman53/project-sprint-halo-suster/helpers"
)

type fileServiceImpl struct {
}

func NewFileService() FileService {
	return &fileServiceImpl{}
}

func (service *fileServiceImpl) Upload(ctx *fiber.Ctx) (file_entity.UploadImageResponse, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return file_entity.UploadImageResponse{}, err
	}
	allowedTypes := []string{"image/jpeg", "image/jpg"}
	if !slices.Contains(allowedTypes, file.Header.Get("Content-Type")) {
		// Handle invalid file type
		return file_entity.UploadImageResponse{}, exc.BadRequestException("Invalid File Type")
	}

	MIN_SIZE, MAX_SIZE := 10*1024, 2*1024*1024
	if file.Size < int64(MIN_SIZE) || file.Size > int64(MAX_SIZE) {
		// Handle invalid file size
		return file_entity.UploadImageResponse{}, exc.BadRequestException("Invalid File Size")
	}

	uniqueId := uuid.New().String()
	extensionName := strings.Split(file.Header.Get("Content-Type"), "/")[1]
	fileName := fmt.Sprintf("%s.%s", uniqueId, extensionName)
	url, err := helpers.FileUpload(file, fileName)
	if err != nil {
		return file_entity.UploadImageResponse{}, err
	}
	return file_entity.UploadImageResponse{
		Message: "Successfully upload image",
		Data:    map[string]string{"imageUrl": url},
	}, nil
}
