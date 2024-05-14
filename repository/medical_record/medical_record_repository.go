package medical_record_repository

import (
	"context"

	medical_record_entity "github.com/kangman53/project-sprint-halo-suster/entity/medical_record"
)

type MedicalRecordRepository interface {
	CreatePatient(ctx context.Context, req medical_record_entity.CreateMRPatientRequest) (medical_record_entity.MRPatientData, error)
}
