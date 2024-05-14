package medical_record_entity

import user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"

type CreatePatientResponse struct {
	Message string       `json:"message"`
	Data    *PatientData `json:"data"`
}

type PatientData struct {
	Id                  string `json:"id,omitempty"`
	IdentityNumber      int    `json:"identityNumber,omitempty"`
	PhoneNumber         string `json:"PhoneNumber,omitempty"`
	Name                string `json:"name,omitempty"`
	BirthDate           string `json:"birthDate,omitempty"`
	Gender              string `json:"gender,omitempty"`
	IdentityCardScanImg string `json:"identityCardScanImg,omitempty"`
	CreatedAt           string `json:"createdAt,omitempty"`
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
	Id             string `json:"id,omitempty"`
	IdentityNumber int    `json:"identityNumber,omitempty"`
	Symptoms       string `json:"symptoms,omitempty"`
	Medications    string `json:"medications,omitempty"`
	CreateAt       string `json:"createdAt,omitempty"`
}

type SearchMedicalRecordResponse struct {
	Message string                     `json:"message"`
	Data    *[]SearchMedicalRecordData `json:"data"`
}

type SearchMedicalRecordData struct {
	IdentityDetail *PatientData          `json:"identityDetail"`
	Symptoms       string                `json:"symptoms"`
	Medications    string                `json:"medications"`
	CreatedAt      string                `json:"createdAt"`
	CreatedBy      *user_entity.UserData `json:"createdBy"`
}
