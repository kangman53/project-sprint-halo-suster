package file_entity

import "mime/multipart"

type UploadImage struct {
	File *multipart.FileHeader
	Size int64
}
