package medical_record_entity

type Patient struct {
	Id                  string
	IdentityNumber      int
	PhoneNumber         string
	Name                string
	BirthDate           string
	Gender              string
	IdentityCardScanImg string
}

type MedicalRecord struct {
	Id             string
	IdentityNumber string
	Symptoms       string
	Name           string
}
