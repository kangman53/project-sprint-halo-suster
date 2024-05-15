package user_entity

type UserResponse struct {
	Message string    `json:"message"`
	Data    *UserData `json:"data"`
}

type UserData struct {
	Id          string `json:"userId"`
	Name        string `json:"name"`
	Nip         int    `json:"nip"`
	AccessToken string `json:"accessToken,omitempty"`
}

type UserResponseData struct {
	Id        string `json:"userId"`
	Name      string `json:"name"`
	Nip       int    `json:"nip"`
	CreatedAt string `json:"createdAt"`
}

type UserGetResponse struct {
	Message string              `json:"message"`
	Data    *[]UserResponseData `json:"data"`
}
