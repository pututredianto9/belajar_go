package conversionHelper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	validationHelper "belajar/helper/validation"

	"github.com/dustin/go-humanize"
	"golang.org/x/crypto/bcrypt"
)

// Convert String Value Into Int
func StrToInt(data string) (int, error) {
	newData, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}

	return newData, nil
}

// Convert String Value Into Int64
func StrToInt64(data string) (int64, error) {
	newData, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}

	return newData, nil
}

// Convert Phone Number Into +62 Format
func ConvertPhoneNumber(data string) string {
	if data[0:1] == "0" {
		data = "+62" + data[1:]
	}
	if data[:2] == "62" {
		data = "+" + data
	}

	return data
}

// Convert Time Into TZ
func ConvertTimeToTZ(data time.Time, timezone string) time.Time {
	loc, _ := time.LoadLocation(timezone)

	newData := data.In(loc)

	return newData
}

// Hide First Some Digit Data
func HideFirstSomeDigitData(data string, digitToShow int) string {
	hideAccountNumber := ""
	for i, v := range data {
		if i < len(data)-digitToShow {
			hideAccountNumber += "*"
			continue
		}

		hideAccountNumber += fmt.Sprintf("%c", v)
	}

	return hideAccountNumber
}

// Convert To RP Currency
func ConvertToRpCurrency(data int64) string {
	humanizeValue := humanize.CommafWithDigits(float64(data), 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)

	return stringValue
}

// Split string use "-"
func SplitStringDash(data string) []string {
	newArray := strings.Split(data, " - ")

	return newArray
}

func ConvertFormatCodesID(codes string, nextID int64) string {
	now := time.Now()

	codes = strings.ReplaceAll(codes, "{{.yyyy}}", now.Format("2006"))
	codes = strings.ReplaceAll(codes, "{{.Month}}", now.Format("January"))
	codes = strings.ReplaceAll(codes, "{{.Mon}}", now.Format("Jan"))
	codes = strings.ReplaceAll(codes, "{{.Day}}", now.Format("Monday"))
	codes = strings.ReplaceAll(codes, "{{.Da}}", now.Format("Mon"))
	codes = strings.ReplaceAll(codes, "{{.dd}}", now.Format("02"))
	codes = strings.ReplaceAll(codes, "{{.mm}}", now.Format("01"))
	codes = strings.ReplaceAll(codes, "{{.yy}}", now.Format("06"))
	codes = strings.ReplaceAll(codes, "{{.d}}", now.Format("2"))
	codes = strings.ReplaceAll(codes, "{{.m}}", now.Format("1"))
	codes = strings.ReplaceAll(codes, "{{.y}}", now.Format("6"))
	codes = strings.ReplaceAll(codes, "{{.id}}", fmt.Sprintf("%06d", nextID))

	return codes
}

func ConvertToPositiveInt(data int) int {
	if data > 0 {
		return data
	}

	return data * -1
}

func ConvertToPositiveFloat64(data float64) float64 {
	if data > 0 {
		return data
	}

	return data * -1
}

func ConvertStringToPlate(data string) string {
	plateSplit := strings.Split(data, "")
	newStr := plateSplit[0]
	for i := 1; i < len(plateSplit); i++ {
		if validationHelper.IsDigit(plateSplit[i-1]) == validationHelper.IsDigit(plateSplit[i]) {
			newStr += plateSplit[i]
			continue
		}

		newStr += " " + plateSplit[i]
	}

	return newStr
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UpperCaseStringBank(a string) string {
	var result string

	upperString := strings.ToUpper(a)

	if len(a) == 3 {
		return strings.ToUpper(a)
	}

	result = strings.Title(strings.ToLower(upperString))
	return result
}

func CheckString(a string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9\s]+$`).MatchString(a)
}

func ReplaceNameAsterisk(a string) string {
	split := strings.Split(a, " ")
	news := ""
	var data []string
	if len(split) > 1 {
		for i, sm := range split {
			newStrings := ""
			if i > 0 {
				for j := 0; j < len(sm); j++ {
					if len(sm) > 2 {
						if j > 0 && j < len(sm)-1 {
							newStrings += "*"
						} else {
							newStrings += string(sm[j])
						}
					} else {
						if j > 0 {
							newStrings += "*"
						} else {
							newStrings += string(sm[j])
						}
					}
				}
				data = append(data, newStrings)
			} else {
				data = append(data, split[0])
			}
		}
		news = strings.Join(data, " ")
	}

	return news
}
