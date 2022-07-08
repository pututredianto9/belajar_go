package paginationHelper

type Page struct {
	Page       int `json:"page"`
	PerPage    int `json:"page_size"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
	First      int `json:"-"`
	Last       int `json:"-"`
}
