package user_entity

type UserRegisterRequest struct {
	Name                string `json:"name" validate:"required,min=5,max=50"`
	Nip                 int    `json:"nip" validate:"required"`
	Password            string `json:"password" validate:"-"`
	Role                string `json:"role" validate:"-"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"-"`
}

type UserLoginRequest struct {
	Nip      int    `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}
