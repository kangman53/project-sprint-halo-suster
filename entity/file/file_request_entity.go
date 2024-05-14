package file_entity

import "mime/multipart"

type UploadImageRequest struct {
	File *multipart.FileHeader
	Name string
}
