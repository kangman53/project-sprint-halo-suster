package medical_record_entity

type CreateMRPatientRequest struct {
	IdentityNumber      int    `json:"identityNumber" validate:"required,int16length"`
	PhoneNumber         string `json:"PhoneNumber" validate:"required,min=10,max=15,phoneNumber"`
	Name                string `json:"name" validate:"required,min=3,max=30"`
	BirthDate           string `json:"birthDate" validate:"required,ISO8601DateTime"`
	Gender              string `json:"gender" validate:"required,gender"`
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,validateUrl"`
}
