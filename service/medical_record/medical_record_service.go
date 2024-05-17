package medical_record_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	medical_record_entity "github.com/kangman53/project-sprint-halo-suster/entity/medical_record"
)

type MedicalRecordService interface {
	CreatePatient(ctx context.Context, req medical_record_entity.CreatePatientRequest) (medical_record_entity.CreatePatientResponse, error)
	SearchPatient(ctx context.Context, query medical_record_entity.SearchPatientQuery) (medical_record_entity.SearchPatientResponse, error)
	CreateMedicalRecord(ctx *fiber.Ctx, req medical_record_entity.CreateMedicalRecordRequest) (medical_record_entity.CreateMedicalRecordResponse, error)
	SearchMedicalRecord(ctx context.Context, query medical_record_entity.SearchMedicalRecordQuery) (medical_record_entity.SearchMedicalRecordResponse, error)
}
