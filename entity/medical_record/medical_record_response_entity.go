package medical_record_entity

type MRPatientData struct {
	Id                  string `json:"id"`
	IdentityNumber      int    `json:"identityNumber,omitempty"`
	PhoneNumber         string `json:"PhoneNumber,omitempty"`
	Name                string `json:"name,omitempty"`
	BirthDate           string `json:"birthDate,omitempty"`
	Gender              string `json:"gender,omitempty"`
	IdentityCardScanImg string `json:"identityCardScanImg,omitempty"`
	CreatedAt           string `json:"createdAt"`
}

type CreateMRPatientResponse struct {
	Message string         `json:"message"`
	Data    *MRPatientData `json:"data"`
}

type SearchMRPatientResponse struct {
	Message string                 `json:"message"`
	Data    *[]MRPatientSearchData `json:"data"`
}

type MRPatientSearchData struct {
	IdentityNumber int    `json:"identityNumber"`
	PhoneNumber    string `json:"PhoneNumber"`
	Name           string `json:"name"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	CreatedAt      string `json:"createdAt"`
}
