package medical_record_repository

import (
	"context"

	medical_record_entity "github.com/kangman53/project-sprint-halo-suster/entity/medical_record"
)

type MedicalRecordRepository interface {
	CreatePatient(ctx context.Context, req medical_record_entity.Patient) (medical_record_entity.Patient, error)
	SearchPatient(ctx context.Context, query medical_record_entity.SearchPatientQuery) (*[]medical_record_entity.SearchPatientData, error)
	CreateMedicalRecord(ctx context.Context, req medical_record_entity.MedicalRecord) (medical_record_entity.MedicalRecord, error)
	SearchMedicalRecord(ctx context.Context, query medical_record_entity.SearchMedicalRecordQuery) (*[]medical_record_entity.SearchMedicalRecordData, error)
}
