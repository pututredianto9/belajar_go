package paginationHelper

import "errors"

var (
	ErrorMaxPage     = errors.New("6001")
	ErrorPage        = errors.New("6002")
	ErrorPageEmpty   = errors.New("6003")
	ErrorPageInvalid = errors.New("6004")
)
