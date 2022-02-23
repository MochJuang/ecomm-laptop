package helper

import (
	"math"
)

type Paginate struct {
	Page       int
	PerPage    int
	PrePage    int
	NextPage   int
	TotalData  int
	TotalPages int
	Data       interface{}
}

func PaginationHelper(items interface{}, page int, perPage int, totalData int) Paginate {
	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = 15
	}

	paginateItems := items
	totalPages := math.Ceil(float64(totalData) / float64(perPage))

	prePage := page - 1
	if prePage < 1 {
		prePage = 0
	}

	var nextPage int
	if int(totalPages) > page {
		nextPage = page + 1
	} else {
		nextPage = 0
	}

	return Paginate{
		Page:       page,
		PerPage:    perPage,
		PrePage:    prePage,
		NextPage:   nextPage,
		TotalData:  totalData,
		TotalPages: int(totalPages),
		Data:       paginateItems,
	}
}
