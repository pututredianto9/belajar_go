package jsonHelper

import "net/http"

//Success response based on excel documentation
const (
	Success   SuccessResponseCode = 1000
	Create    SuccessResponseCode = 1001
	NoData    SuccessResponseCode = 1004
	NoUpdated SuccessResponseCode = 1005
)

type SuccessResponseCode int

var successText = map[SuccessResponseCode]string{
	Success:   "OK",
	NoData:    "There is no data to send for this request",
	NoUpdated: "There is No updated data",
}

func SuccessText(code SuccessResponseCode) string {
	return successText[code]
}

var httpStatus = map[SuccessResponseCode]int{
	Success:   http.StatusOK,
	Create:    http.StatusCreated,
	NoData:    http.StatusOK,
	NoUpdated: http.StatusOK,
}

func HttpStatusCode(code SuccessResponseCode) int {
	return httpStatus[code]
}
