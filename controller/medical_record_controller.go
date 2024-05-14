package controller

import (
	"github.com/gofiber/fiber/v2"
	medical_record_entity "github.com/kangman53/project-sprint-halo-suster/entity/medical_record"
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	medical_record_service "github.com/kangman53/project-sprint-halo-suster/service/medical_record"
)

type MedicalRecordController struct {
	MedicalRecordService medical_record_service.MedicalRecordService
}

func NewMedicalRecordController(medicalRecordService medical_record_service.MedicalRecordService) *MedicalRecordController {
	return &MedicalRecordController{
		MedicalRecordService: medicalRecordService,
	}
}

func (controller *MedicalRecordController) CreatePatient(ctx *fiber.Ctx) error {
	patientReq := new(medical_record_entity.CreateMRPatientRequest)
	if err := ctx.BodyParser(patientReq); err != nil {
		return exc.BadRequestException("Failed to parse request body")
	}
	resp, err := controller.MedicalRecordService.CreatePatient(ctx.UserContext(), *patientReq)
	if err != nil {
		return exc.Exception(ctx, err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (controller *MedicalRecordController) SearchPatient(ctx *fiber.Ctx) error {
	searchQuery := new(medical_record_entity.SearchMRPatientQuery)
	searchQuery.Limit = 5
	searchQuery.Offset = 0

	if err := ctx.QueryParser(searchQuery); err != nil {
		return exc.BadRequestException("Error when parsing request query")
	}

	resp, err := controller.MedicalRecordService.SearchPatient(ctx.UserContext(), *searchQuery)
	if err != nil {
		return exc.Exception(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (controller *MedicalRecordController) CreateMedicalRecord(ctx *fiber.Ctx) error {
	medicalRecord := new(medical_record_entity.CreateMedicalRecordRequest)
	if err := ctx.BodyParser(medicalRecord); err != nil {
		return exc.BadRequestException("Failed to parse request body")
	}
	resp, err := controller.MedicalRecordService.CreateMedicalRecord(ctx, *medicalRecord)
	if err != nil {
		return exc.Exception(ctx, err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}
