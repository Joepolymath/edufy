package auth

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// custom  password validator to be regsitered in any validator instance needed
func ValidateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Define your password criteria using a regular expression
	regex := regexp.MustCompile(`[a-z]`)
	if !regex.MatchString(password) {
		return false
	}

	regex = regexp.MustCompile(`[A-Z]`)
	if !regex.MatchString(password) {
		return false
	}

	regex = regexp.MustCompile(`\d`)
	if !regex.MatchString(password) {
		return false
	}

	regex = regexp.MustCompile(`[^a-zA-Z\d]`)
	if !regex.MatchString(password) {
		return false
	}

	return true
}
