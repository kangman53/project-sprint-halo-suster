package medical_record_repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
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
	RETURNING id, to_char(created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') created_at`
	if err := repostory.DBpool.QueryRow(ctx, query, identityNumber, patient.PhoneNumber, patient.Name, patient.Gender, patient.BirthDate, patient.IdentityCardScanImg).Scan(&patientId, &createdAt); err != nil {
		return medical_record_entity.MRPatientData{}, err
	}

	return medical_record_entity.MRPatientData{Id: patientId, CreatedAt: createdAt}, nil
}

func (repository *medicalRecordRepositoryImpl) SearchPatient(ctx context.Context, searchQuery medical_record_entity.SearchMRPatientQuery) (*[]medical_record_entity.MRPatientSearchData, error) {
	query := `SELECT CAST(identity_number AS BIGINT), phone_number, name, to_char(birth_date, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') birth_date, gender, to_char(created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') created_at
	FROM patients`

	var whereClause []string
	var searchParams []interface{}

	if searchQuery.IdentityNumber != "" {
		whereClause = append(whereClause, fmt.Sprintf("identity_number = $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.IdentityNumber)
	}
	if searchQuery.Name != "" {
		whereClause = append(whereClause, fmt.Sprintf("name ~* $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.Name)
	}
	if searchQuery.PhoneNumber != "" {
		whereClause = append(whereClause, fmt.Sprintf("phone_number ~* $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.PhoneNumber)
	}

	if len(whereClause) > 0 {
		query += " WHERE " + strings.Join(whereClause, " AND ")
	}

	var orderBy string
	if strings.ToLower(searchQuery.CreatedAt) == "asc" {
		orderBy = ` ORDER BY created_at ASC`
	} else {
		orderBy = ` ORDER BY created_at DESC`
	}
	query += orderBy

	// handle limit or offset negative
	if searchQuery.Limit < 0 {
		searchQuery.Limit = 5
	}
	if searchQuery.Offset < 0 {
		searchQuery.Offset = 0
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", searchQuery.Limit, searchQuery.Offset)
	rows, err := repository.DBpool.Query(ctx, query, searchParams...)
	if err != nil {
		return &[]medical_record_entity.MRPatientSearchData{}, err
	}
	defer rows.Close()

	patients, err := pgx.CollectRows(rows, pgx.RowToStructByName[medical_record_entity.MRPatientSearchData])
	if err != nil {
		return &[]medical_record_entity.MRPatientSearchData{}, err
	}

	return &patients, nil
}
