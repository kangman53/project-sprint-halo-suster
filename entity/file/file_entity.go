package file_entity

type UploadImageResponse struct {
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}
