package medical_record_service

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	medical_record_entity "github.com/kangman53/project-sprint-halo-suster/entity/medical_record"
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	medical_record_repository "github.com/kangman53/project-sprint-halo-suster/repository/medical_record"
	authService "github.com/kangman53/project-sprint-halo-suster/service/auth"
)

type medicalRecordServiceImpl struct {
	MedicalRecordRepository medical_record_repository.MedicalRecordRepository
	AuthService             authService.AuthService
	Validator               *validator.Validate
}

func NewMedicalRecordService(medicalRecordRepostory medical_record_repository.MedicalRecordRepository, authService authService.AuthService, validator *validator.Validate) MedicalRecordService {
	return &medicalRecordServiceImpl{
		MedicalRecordRepository: medicalRecordRepostory,
		AuthService:             authService,
		Validator:               validator,
	}
}

func (service *medicalRecordServiceImpl) CreatePatient(ctx context.Context, req medical_record_entity.CreatePatientRequest) (medical_record_entity.CreatePatientResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(req); err != nil {
		return medical_record_entity.CreatePatientResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
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
			return medical_record_entity.CreatePatientResponse{}, exc.ConflictException("Patient with this identity number already registered")
		}
		return medical_record_entity.CreatePatientResponse{}, err
	}

	return medical_record_entity.CreatePatientResponse{
		Message: "Patient successfully created",
		Data:    &patientCreated,
	}, nil
}

func (service *medicalRecordServiceImpl) SearchPatient(ctx context.Context, searchQuery medical_record_entity.SearchPatientQuery) (medical_record_entity.SearchPatientResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(searchQuery); err != nil {
		return medical_record_entity.SearchPatientResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	patientSearched, err := service.MedicalRecordRepository.SearchPatient(ctx, searchQuery)
	if err != nil {
		return medical_record_entity.SearchPatientResponse{}, err
	}

	return medical_record_entity.SearchPatientResponse{
		Message: "success",
		Data:    patientSearched,
	}, nil
}

func (service *medicalRecordServiceImpl) CreateMedicalRecord(ctx *fiber.Ctx, req medical_record_entity.CreateMedicalRecordRequest) (medical_record_entity.CreateMedicalRecordResponse, error) {
	if err := service.Validator.Struct(req); err != nil {
		return medical_record_entity.CreateMedicalRecordResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	userId, err := service.AuthService.GetValidUser(ctx)
	if err != nil {
		return medical_record_entity.CreateMedicalRecordResponse{}, exc.UnauthorizedException("Unauthorized")
	}

	medicalRecord := medical_record_entity.MedicalRecord{
		IdentityNumber: req.IdentityNumber,
		Symptoms:       req.Symptoms,
		Medications:    req.Medications,
		CreatedBy:      userId,
	}
	medicalRecordCreated, err := service.MedicalRecordRepository.CreateMedicalRecord(ctx.UserContext(), medicalRecord)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return medical_record_entity.CreateMedicalRecordResponse{}, exc.BadRequestException("identityNumber is not exist")
		}
		return medical_record_entity.CreateMedicalRecordResponse{}, err
	}

	return medical_record_entity.CreateMedicalRecordResponse{
		Message: "Medical Record successfully created",
		Data:    &medicalRecordCreated,
	}, nil
}

func (service *medicalRecordServiceImpl) SearchMedicalRecord(ctx context.Context, searchQuery medical_record_entity.SearchMedicalRecordQuery) (medical_record_entity.SearchMedicalRecordResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(searchQuery); err != nil {
		return medical_record_entity.SearchMedicalRecordResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	medicalRecordSearched, err := service.MedicalRecordRepository.SearchMedicalRecord(ctx, searchQuery)
	if err != nil {
		return medical_record_entity.SearchMedicalRecordResponse{}, err
	}

	return medical_record_entity.SearchMedicalRecordResponse{
		Message: "success",
		Data:    medicalRecordSearched,
	}, nil
}
