package medical_record_entity

import user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"

type CreatePatientResponse struct {
	Message string   `json:"message"`
	Data    *Patient `json:"data"`
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
	Message string         `json:"message"`
	Data    *MedicalRecord `json:"data"`
}

type SearchMedicalRecordResponse struct {
	Message string                     `json:"message"`
	Data    *[]SearchMedicalRecordData `json:"data"`
}

type SearchMedicalRecordData struct {
	IdentityDetail *Patient              `json:"identityDetail"`
	Symptoms       string                `json:"symptoms"`
	Medications    string                `json:"medications"`
	CreatedAt      string                `json:"createdAt"`
	CreatedBy      *user_entity.UserData `json:"createdBy"`
}
