package jsonHelper

import (
	"encoding/json"
	"net/http"
	"strconv"

	errorHelper "belajar/helper/error"
	loggingHelper "belajar/helper/log"
	paginationHelper "belajar/helper/pagination"
)

// Return JSON
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// Return JSON Error
func ErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message interface{}) {
	res := &response{
		Error: message,
	}

	loggingHelper.Addlog(r, "ERROR", message)

	WriteJSON(w, statusCode, res)
}

// Return JSON Success V3
func SuccessResponseV3(w http.ResponseWriter, r *http.Request, status bool, code string, data interface{}, pagination *paginationHelper.Page) {
	meta := &meta{
		Status:     status,
		Message:    "success",
		Code:       "OK",
		Pagination: pagination,
	}

	res := &responseV3{
		Meta: *meta,
		Data: data,
	}

	loggingHelper.Addlog(r, "SUCCESS", data)

	WriteJSON(w, http.StatusOK, res)
}

// Return JSON Success V3 with HTTP Code
func SuccessResponseV3byHttpCode(w http.ResponseWriter, r *http.Request, status bool, httpStatus int, code string, data interface{}, message string, pagination *paginationHelper.Page) {
	meta := &meta{
		Status:     status,
		Message:    message,
		Code:       code,
		Pagination: pagination,
	}

	res := &responseV3{
		Meta: *meta,
		Data: data,
	}

	loggingHelper.Addlog(r, "SUCCESS", data)

	WriteJSON(w, httpStatus, res)
}

// Return JSON Success V4
func SuccessResponseV4(w http.ResponseWriter, r *http.Request, status bool, code SuccessResponseCode, data interface{}, pagination *paginationHelper.Page) {
	meta := &meta{
		Status:     status,
		Message:    SuccessText(code),
		Code:       strconv.Itoa(int(code)),
		Pagination: pagination,
	}

	res := &responseV3{
		Meta: *meta,
		Data: data,
	}

	loggingHelper.Addlog(r, "SUCCESS", res)

	WriteJSON(w, HttpStatusCode(code), res)
}

// Return JSON Error V3
func ErrorResponseV3(w http.ResponseWriter, r *http.Request, status bool, code string, message interface{}) {
	meta := &meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	res := &responseV3{
		Meta: *meta,
		Data: nil,
	}

	loggingHelper.Addlog(r, "ERROR", message)

	WriteJSON(w, http.StatusBadRequest, res)
}

// Return JSON Error V4
func ErrorResponseV4(w http.ResponseWriter, r *http.Request, status bool, httpCode int, message errorHelper.Message) {

	meta := &meta{
		Code:    message.Code,
		Status:  status,
		Message: message.MessageBE,
	}

	res := &responseV3{
		Meta: *meta,
		Data: nil,
	}

	loggingHelper.Addlog(r, "ERROR", res)

	WriteJSON(w, httpCode, res)
}

// Return JSON Error V4 with Data
func ErrorResponseV4WithData(w http.ResponseWriter, r *http.Request, status bool, httpCode int, message errorHelper.Message, data interface{}) {

	meta := &meta{
		Code:    message.Code,
		Status:  status,
		Message: message.MessageBE,
	}

	res := &responseV3{
		Meta: *meta,
		Data: data,
	}

	loggingHelper.Addlog(r, "ERROR", res)

	WriteJSON(w, httpCode, res)
}
