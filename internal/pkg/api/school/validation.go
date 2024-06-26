package school

import (
	"Learnium/internal/pkg/models"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateSchoolType(fl validator.FieldLevel) bool {
	schoolType := fl.Field().String()
	// Convert input to uppercase for case-insensitive comparison
	return ValidateRoleTypeByInput(schoolType)
}

func ValidateSchoolTypeByInput(input string) bool {
	input = strings.ToUpper(input)

	// Check if the input matches any of the model types
	switch models.SchoolType(input) {
	case models.K12, models.RELIGIOUS, models.TERTIARY, models.SUMMER, models.GOV, models.VOCATIONAL:
		return true
	default:
		return false
	}
}

func ValidateRoleTypeByInput(input string) bool {
	input = strings.ToUpper(input)

	// Check if the input matches any of the model types
	switch models.RoleType(input) {
	case models.ADMIN, models.TEACHING_STAFF, models.NON_TEACHING_STAFF:
		return true
	default:
		return false
	}
}
