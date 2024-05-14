package medical_record_repository

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	medical_record_entity "github.com/kangman53/project-sprint-halo-suster/entity/medical_record"
)

type medicalRecordRepositoryImpl struct {
	DBpool *pgxpool.Pool
}

func NewMedicalRecordRepository(dbPool *pgxpool.Pool) MedicalRecordRepository {
	return &medicalRecordRepositoryImpl{
		DBpool: dbPool,
	}
}

func (repostory *medicalRecordRepositoryImpl) CreatePatient(ctx context.Context, patient medical_record_entity.CreateMRPatientRequest) (medical_record_entity.MRPatientData, error) {
	var identityNumber, patientId, createdAt string
	identityNumber = strconv.Itoa(patient.IdentityNumber)

	query := `INSERT INTO patients (identity_number, phone_number, name, gender, birth_date, identity_card_scan_img) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, to_char(created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') createdAt`
	if err := repostory.DBpool.QueryRow(ctx, query, identityNumber, patient.PhoneNumber, patient.Name, patient.Gender, patient.BirthDate, patient.IdentityCardScanImg).Scan(&patientId, &createdAt); err != nil {
		return medical_record_entity.MRPatientData{}, err
	}

	return medical_record_entity.MRPatientData{Id: patientId, CreatedAt: createdAt}, nil
}
