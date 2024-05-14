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

func (repository *medicalRecordRepositoryImpl) CreatePatient(ctx context.Context, patient medical_record_entity.Patient) (medical_record_entity.PatientData, error) {
	var identityNumber, patientId, createdAt string
	identityNumber = strconv.Itoa(patient.IdentityNumber)

	query := `INSERT INTO patients (identity_number, phone_number, name, gender, birth_date, identity_card_scan_img) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, to_char(created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') created_at`
	if err := repository.DBpool.QueryRow(ctx, query, identityNumber, patient.PhoneNumber, patient.Name, patient.Gender, patient.BirthDate, patient.IdentityCardScanImg).Scan(&patientId, &createdAt); err != nil {
		return medical_record_entity.PatientData{}, err
	}

	return medical_record_entity.PatientData{Id: patientId, CreatedAt: createdAt}, nil
}

func (repository *medicalRecordRepositoryImpl) SearchPatient(ctx context.Context, searchQuery medical_record_entity.SearchPatientQuery) (*[]medical_record_entity.SearchPatientData, error) {
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
		return &[]medical_record_entity.SearchPatientData{}, err
	}
	defer rows.Close()

	patients, err := pgx.CollectRows(rows, pgx.RowToStructByName[medical_record_entity.SearchPatientData])
	if err != nil {
		return &[]medical_record_entity.SearchPatientData{}, err
	}

	return &patients, nil
}

func (repository *medicalRecordRepositoryImpl) CreateMedicalRecord(ctx context.Context, req medical_record_entity.MedicalRecord) (medical_record_entity.MedicalRecordData, error) {
	var identityNumber, medicalId, createdAt string
	identityNumber = strconv.Itoa(req.IdentityNumber)

	query := `INSERT INTO medical_records (patient_identity_number, symptoms, medications, created_by) 
	SELECT 
		$1, $2, $3, $4
	WHERE EXISTS (
		SELECT 1 FROM patients WHERE identity_number = $5
	)
	RETURNING id, to_char(created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') created_at;`
	if err := repository.DBpool.QueryRow(ctx, query, identityNumber, req.Symptoms, req.Medications, req.CreateBy, identityNumber).Scan(&medicalId, &createdAt); err != nil {
		return medical_record_entity.MedicalRecordData{}, err
	}

	return medical_record_entity.MedicalRecordData{Id: medicalId, CreateAt: createdAt}, nil
}

func (repository *medicalRecordRepositoryImpl) SearchMedicalRecord(ctx context.Context, searchQuery medical_record_entity.SearchMedicalRecordQuery) (*[]medical_record_entity.SearchMedicalRecordData, error) {
	query := `SELECT 
	json_build_object('identityNumber', CAST(p.identity_number AS BIGINT), 'phoneNumber', p.phone_number, 'name', p.name, 'birthDate', to_char(p.birth_date, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'), 'gender', p.gender, 'identityCardScanImg', p.identity_card_scan_img) identity_detail,
	m.symptoms,
	m.medications,
	to_char(m.created_at, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') created_at,
	json_build_object('nip', CAST(u.nip AS BIGINT), 'name', u.name, 'userId', u.id) created_by
	FROM medical_records m
		JOIN users u ON u.id = m.created_by
		JOIN patients p ON p.identity_number = m.patient_identity_number
	`

	var whereClause []string
	var searchParams []interface{}

	if searchQuery.IdentityNumber != "" {
		whereClause = append(whereClause, fmt.Sprintf("m.patient_identity_number = $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.IdentityNumber)
	}
	if searchQuery.CreatedById != "" {
		whereClause = append(whereClause, fmt.Sprintf("m.created_by = $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.CreatedById)
	}
	if searchQuery.CreatedByNip != "" {
		whereClause = append(whereClause, fmt.Sprintf("u.nip = $%d", len(searchParams)+1))
		searchParams = append(searchParams, searchQuery.CreatedByNip)
	}

	if len(whereClause) > 0 {
		query += " WHERE " + strings.Join(whereClause, " AND ")
	}

	var orderBy string
	if strings.ToLower(searchQuery.CreatedAt) == "asc" {
		orderBy = ` ORDER BY m.created_at ASC`
	} else {
		orderBy = ` ORDER BY m.created_at DESC`
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
		return &[]medical_record_entity.SearchMedicalRecordData{}, err
	}
	defer rows.Close()

	medicalRecords, err := pgx.CollectRows(rows, pgx.RowToStructByName[medical_record_entity.SearchMedicalRecordData])
	if err != nil {
		return &[]medical_record_entity.SearchMedicalRecordData{}, err
	}

	return &medicalRecords, nil
}
