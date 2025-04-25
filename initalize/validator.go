package initalize

import (
	myvalidator "my-ops-admin/pkg/validator"

	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	validate := validator.New()
	validate.RegisterValidation("dateformat", myvalidator.ValidateDateFormat)
}
