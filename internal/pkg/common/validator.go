package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "Learnium/internal"
)

// validators are useful for maintaining data compliance ensuring that all incoming data particularly
// over http requests meet predefined criteria and thus fit for purpose

// describe behaviour of validator object
type IValidator interface {
	ValidateRequestBody(requestContext *fiber.Ctx, model interface{}) error
	RegisterCustomValidator(name string, fn validator.Func) error
}

// model a validator object
type Validator struct {
	validator *validator.Validate
}

// create new validator instance
func NewValidator() IValidator {
	return &Validator{
		validator.New(),
	}
	// return newValidator
}

// register any custom validator function of your choice
func (v *Validator) RegisterCustomValidator(name string, fn validator.Func) error {
	return v.validator.RegisterValidation(name, fn)
}

func (v *Validator) ValidateRequestBody(requestContext *fiber.Ctx, model interface{}) error {

	if err := requestContext.BodyParser(&model); err != nil {
		return err
	}
	/* this is used to validate models in which we used instead of the creation serializer*/

	if err := v.validator.Struct(model); err != nil {
		// Cast the error to validator.ValidationErrors to access the actual errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// errorMessages := make(map[string]string)
			result := "{"
			for _, err := range validationErrors {
				errorMessage := strings.Split(err.Error(), err.Field()+"'")
				result += fmt.Sprintf(`"%s": "%v", `, err.Field(), errorMessage[2])
			}

			result = result[:len(result)-2] + "}"

			return errors.New(result)

		}
		return errors.New("Request Data Invalid")
	}

	return nil
}

func ValidateUUIDByInput(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}
