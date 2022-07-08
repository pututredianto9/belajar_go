package errorHelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ErrorDataNotExist(object, identifier string) error {
	return fmt.Errorf("%s with identifier %s doesn't exists", object, identifier)
}

func ErrorDataUnverified(object, identifier string) error {
	return fmt.Errorf("%s with identifier %s is still under review", object, identifier)
}

func ErrorDataUpdate(object, identifier string) error {
	return fmt.Errorf("failed to update %s with identifier %s", object, identifier)
}

func ErrorDataDelete(object, identifier string) error {
	return fmt.Errorf("failed to delete %s with identifier %s", object, identifier)
}

func ErrorDataCannotEmpty(object string) error {
	return fmt.Errorf("%s field cannot be empty", object)
}

func ErrorDataCannotLessThenZero(object string) error {
	return fmt.Errorf("%s field cannot less then zero", object)
}

func ErrorDataCannotAfterNow() error {
	return fmt.Errorf("sorry, to checkout your cart delivery date must be at least 1 day from now")
}

func ErrorWrongData(object string) error {
	return fmt.Errorf("invalid input %s", object)
}

func ErrorPhotoSize(object string) error {
	return fmt.Errorf("%s photo too large", object)
}

func ErrorPhotoType(object string) error {
	return fmt.Errorf("%s make sure photo format is png, jpg, jpeg", object)
}

func ErrorDataExist(object string) error {
	return fmt.Errorf("%s has been registered", object)
}

type ErrorMessage interface {
	GetMessage(string) Message
	InitJSON(string) error
}

type JsonError struct {
	ErrMessage map[string]Message
}

func NewErrorMessage() ErrorMessage {
	json := new(JsonError)

	return json
}

func (c *JsonError) InitJSON(file string) error {
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()

	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]Message

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		return err
	}

	c.ErrMessage = result
	return nil
}

func (c *JsonError) GetMessage(code string) Message {
	num := c.ErrMessage[code]
	num.Code = code
	return num
}
