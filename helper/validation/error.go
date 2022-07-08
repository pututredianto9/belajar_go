package validationHelper

import (
	"errors"
)

var (
	ErrorStrIsNotDigit = errors.New("string is not digit")
)

var (
	ErrorInvalidPhoneNumber = errors.New("invalid phone number format")
)
