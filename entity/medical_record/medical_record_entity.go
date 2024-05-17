package medical_record_entity

type Patient struct {
	Id                  string `json:"id,omitempty"`
	IdentityNumber      int    `json:"identityNumber,omitempty"`
	PhoneNumber         string `json:"PhoneNumber,omitempty"`
	Name                string `json:"name,omitempty"`
	BirthDate           string `json:"birthDate,omitempty"`
	Gender              string `json:"gender,omitempty"`
	IdentityCardScanImg string `json:"identityCardScanImg,omitempty"`
	CreatedAt           string `json:"createdAt,omitempty"`
}

type MedicalRecord struct {
	Id             string `json:"id,omitempty"`
	IdentityNumber int    `json:"identityNumber,omitempty"`
	Symptoms       string `json:"symptoms,omitempty"`
	Medications    string `json:"medications,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	CreateAt       string `json:"createdAt,omitempty"`
}
