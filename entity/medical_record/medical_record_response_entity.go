package medical_record_entity

type PatientData struct {
	Id                  string `json:"id"`
	IdentityNumber      int    `json:"identityNumber,omitempty"`
	PhoneNumber         string `json:"PhoneNumber,omitempty"`
	Name                string `json:"name,omitempty"`
	BirthDate           string `json:"birthDate,omitempty"`
	Gender              string `json:"gender,omitempty"`
	IdentityCardScanImg string `json:"identityCardScanImg,omitempty"`
	CreatedAt           string `json:"createdAt"`
}

type CreatePatientResponse struct {
	Message string       `json:"message"`
	Data    *PatientData `json:"data"`
}

type SearchPatientResponse struct {
	Message string               `json:"message"`
	Data    *[]SearchPatientData `json:"data"`
}

type SearchPatientData struct {
	IdentityNumber int    `json:"identityNumber"`
	PhoneNumber    string `json:"PhoneNumber"`
	Name           string `json:"name"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	CreatedAt      string `json:"createdAt"`
}

type CreateMedicalRecordResponse struct {
	Message string             `json:"message"`
	Data    *MedicalRecordData `json:"data"`
}

type MedicalRecordData struct {
	Id       string `json:"id"`
	CreateAt string `json:"createdAt"`
}
