package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/go-playground/validator"
)

func validatePhoneNumber(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := `^\+62\d+$`
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

func validateUrl(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := `^(?:https?:\/\/)?(?:www\.)?(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(?:\/[^\s]*)?$`
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

func validateProductCategory(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	for _, categ := range ProductCategory {
		if categ == value {
			return true
		}
	}
	return false
}

func validatNipByRole(value string, code string) bool {

	pattern := fmt.Sprintf(`^%s[12]20(?:0[0-9]|1[0-9]|2[0-4])(0[1-9]|1[0-2])([0-9]{3,5})$`, code)
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

func validateITNip(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	return validatNipByRole(value, "615")
}

func validateNurseNip(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	return validatNipByRole(value, "303")
}

func validateGender(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == "male" || value == "female"
}

func validateISO8601DateTime(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		_, err := time.Parse(time.RFC3339, value)
		return err == nil
	}
	return false
}

func validateInt16Length(fl validator.FieldLevel) bool {
	str := strconv.Itoa(fl.Field().Interface().(int))
	return len(str) == 16
}

func RegisterCustomValidator(validator *validator.Validate) {
	// validator.RegisterValidation() -> if you want to create new tags rule to be used on struct entity
	// validator.RegisterStructValidation() -> if you want to create validator then access all fields to the struct entity

	validator.RegisterValidation("phoneNumber", validatePhoneNumber)
	validator.RegisterValidation("productCategory", validateProductCategory)
	validator.RegisterValidation("validateUrl", validateUrl)
	validator.RegisterValidation("nipIT", validateITNip)
	validator.RegisterValidation("nipNurse", validateNurseNip)
	validator.RegisterValidation("gender", validateGender)
	validator.RegisterValidation("ISO8601DateTime", validateISO8601DateTime)
	validator.RegisterValidation("int16length", validateInt16Length)
}
