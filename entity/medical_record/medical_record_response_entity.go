package medical_record_entity

type MRPatientData struct {
	Id                  string `json:"id,omitempty"`
	IdentityNumber      int    `json:"identityNumber,omitempty"`
	PhoneNumber         string `json:"PhoneNumber,omitempty"`
	Name                string `json:"name,omitempty"`
	BirthDate           string `json:"birthDate,omitempty"`
	Gender              string `json:"gender,omitempty"`
	IdentityCardScanImg string `json:"identityCardScanImg,omitempty"`
	CreatedAt           string `json:"createdAt,omitempty"`
}

type CreateMRPatientResponse struct {
	Message string         `json:"message"`
	Data    *MRPatientData `json:"data"`
}
