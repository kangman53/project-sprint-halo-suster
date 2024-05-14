package medical_record_service

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	medical_record_entity "github.com/kangman53/project-sprint-halo-suster/entity/medical_record"
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	medical_record_repository "github.com/kangman53/project-sprint-halo-suster/repository/medical_record"
)

type medicalRecordServiceImpl struct {
	MedicalRecordRepository medical_record_repository.MedicalRecordRepository
	Validator               *validator.Validate
}

func NewMedicalRecordService(medicalRecordRepostory medical_record_repository.MedicalRecordRepository, validator *validator.Validate) MedicalRecordService {
	return &medicalRecordServiceImpl{
		MedicalRecordRepository: medicalRecordRepostory,
		Validator:               validator,
	}
}

func (service *medicalRecordServiceImpl) CreatePatient(ctx context.Context, req medical_record_entity.CreateMRPatientRequest) (medical_record_entity.CreateMRPatientResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(req); err != nil {
		return medical_record_entity.CreateMRPatientResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	patient := medical_record_entity.Patient{
		IdentityNumber:      req.IdentityNumber,
		PhoneNumber:         req.PhoneNumber,
		Name:                req.Name,
		BirthDate:           req.BirthDate,
		Gender:              req.Gender,
		IdentityCardScanImg: req.IdentityCardScanImg,
	}
	patientCreated, err := service.MedicalRecordRepository.CreatePatient(ctx, patient)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return medical_record_entity.CreateMRPatientResponse{}, exc.ConflictException("Patient with this identity number already registered")
		}
		return medical_record_entity.CreateMRPatientResponse{}, err
	}

	return medical_record_entity.CreateMRPatientResponse{
		Message: "Patient successfully created",
		Data:    &patientCreated,
	}, nil
}

func (service *medicalRecordServiceImpl) SearchPatient(ctx context.Context, searchQuery medical_record_entity.SearchMRPatientQuery) (medical_record_entity.SearchMRPatientResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(searchQuery); err != nil {
		return medical_record_entity.SearchMRPatientResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	patientSearched, err := service.MedicalRecordRepository.SearchPatient(ctx, searchQuery)
	if err != nil {
		return medical_record_entity.SearchMRPatientResponse{}, err
	}

	return medical_record_entity.SearchMRPatientResponse{
		Message: "success",
		Data:    patientSearched,
	}, nil
}
