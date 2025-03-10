package controllers

import (
	"math"
	"net/http"
	"strconv"
)

type PaginationData struct {
	CurrentPage int
	TotalPages  int
	TotalItems  int
	Query       string
	Start       int
	End         int
}

func GetPaginationData(r *http.Request, totalItems int, itemsPerPage int) PaginationData {
	query := r.URL.Query().Get("search")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))

	if page > totalPages && totalPages > 0 {
		page = totalPages
	}

	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage
	if end > totalItems {
		end = totalItems
	}

	return PaginationData{
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
		Query:       query,
		Start:       start,
		End:         end,
	}
}
