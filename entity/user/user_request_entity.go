package user_entity

type UserRegisterRequest struct {
	Name                string `json:"name" validate:"required,min=5,max=50"`
	Nip                 int    `json:"nip" validate:"required"`
	Password            string `json:"password" validate:"-"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"-"`
}

type NurseEditRequest struct {
	Name string `json:"name" validate:"required,min=5,max=50"`
	Nip  int    `json:"nip" validate:"required"`
}

type UserLoginRequest struct {
	Nip      int    `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type NurseAccessRequest struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type UserGetRequest struct {
	Id        string `query:"userId"`
	Name      string `query:"name"`
	Nip       string `query:"nip"`
	Role      string `query:"role"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
	CreatedAt string `query:"createdAt"`
}
