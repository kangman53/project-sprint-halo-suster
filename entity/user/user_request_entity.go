package user_entity

type UserRegisterRequest struct {
	Name                string `json:"name" validate:"required,min=5,max=50"`
	Nip                 int    `json:"nip" validate:"required"`
	Password            string `json:"password" validate:"-"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"-"`
}

type NurseEditRequest struct {
	Name                string `json:"name" validate:"required,min=5,max=50"`
	Nip                 int    `json:"nip" validate:"required"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,validateUrl"`
}

type UserLoginRequest struct {
	Nip      int    `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type NurseAccessRequest struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type UserGetRequest struct {
	Id        string `json:"userId"`
	Name      string `json:"name"`
	Nip       int    `json:"nip"`
	Role      string `json:"role"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	CreatedAt string `json:"createdAt"`
}
