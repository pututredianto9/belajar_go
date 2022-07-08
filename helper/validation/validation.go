package validationHelper

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Validate That String Data Is A Digit
func StrIsDigit(data string) error {
	for _, v := range data {
		isDigit := unicode.IsDigit(v)
		if !isDigit {
			return ErrorStrIsNotDigit
		}
	}

	return nil
}

// Validate Phone Number With Indonesian Format
func ValidatePhoneNumber(data string) error {
	r, _ := regexp.Compile(RegexPhoneNumber)

	res := r.MatchString(data)
	if !res {
		return ErrorInvalidPhoneNumber
	}

	return nil
}

// Validate check license plate format
func ValidateLicencePlate(data string) bool {
	re := regexp.MustCompile(RegexLicensePlate)
	plate := re.MatchString(data)
	if plate {
		return true
	}

	return false
}

// Validation check digit
func IsDigit(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	return true
}

// Validation for check value
func ValidateTextValue(i int) string {
	if i >= 0 && i <= 20 {
		return "Perlu ganti sebelum 1,000 km"
	} else if i >= 30 && i <= 40 {
		return "Perlu ganti sebelum 10,000 km"
	} else if i >= 50 && i <= 60 {
		return "Perlu service atau perawatan"
	} else if i >= 70 && i <= 80 {
		return "Aman, tapi butuh cek rutin"
	} else if i >= 90 && i <= 100 {
		return "Prima/baru"
	} else {
		return ""
	}
}

// Validation for check color value inspection
func ValidateColorInspection(i int) string {
	if i >= 0 && i <= 20 {
		return "#ED2E2E"
	} else if i >= 30 && i <= 40 {
		return "#FF6400"
	} else if i >= 50 && i <= 60 {
		return "#EB9C03"
	} else if i >= 70 && i <= 80 {
		return "#00BA88"
	} else if i >= 90 && i <= 100 {
		return "#00966D"
	} else {
		return ""
	}
}
