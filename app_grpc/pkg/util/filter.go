package util

import (
	"regexp"
	"strings"
)

type Pagination struct {
	Limit int32
	Page  int32
	Field string
	Sort  string
}

func GeneratePaginationFromRequest(limit int32, page int32, field string, sort string) Pagination {
	isValidLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if !isValidLetter(field) {
		field = "created_at"
	}

	switch {
	case strings.ToLower(sort) == "asc":
		sort = "ASC"
	case strings.ToLower(sort) == "desc":
		sort = "DESC"
	default:
		sort = "ASC"
	}

	if limit == 0 || page == 0 {
		limit = 2
		page = 1
	}

	return Pagination{
		Limit: limit,
		Page:  page,
		Field: field,
		Sort:  sort,
	}
}
