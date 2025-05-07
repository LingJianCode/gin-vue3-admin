package myvalidator

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

// format: 2006-04-01
func ValidateDateFormat(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) != 10 || value[4] != '-' || value[7] != '-' {
		return false
	}
	yearStr := value[:4]
	monthStr := value[5:7]
	dayStr := value[8:]
	_, errYear := strconv.Atoi(yearStr)
	if errYear != nil {
		return false
	}
	month, errMonth := strconv.Atoi(monthStr)
	if errMonth != nil || month < 1 || month > 12 {
		return false
	}
	day, errDay := strconv.Atoi(dayStr)
	if errDay != nil || day < 1 || day > 31 {
		return false
	}
	return true
}
