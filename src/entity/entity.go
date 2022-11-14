package entity

import (
	"math"
)

type PaginationParam struct {
	Limit          int64 `form:"limit" param:"limit" json:"-"`
	Page           int64 `form:"page" param:"page" json:"-"`
	CurrentPage    int64 `json:"currentPage"`
	TotalPage      int64 `json:"totalPage"`
	CurrentElement int64 `json:"currentElement"`
	TotalElement   int64 `json:"totalElement"`
}

func (pg *PaginationParam) ProcessPagination() {
	pg.CurrentPage = pg.Page
	pg.TotalPage = int64(math.Ceil(float64(pg.TotalElement) / float64(pg.Limit)))
	pg.CurrentElement = pg.Limit
}

func (pg *PaginationParam) GetOffset() int64 {
	return (pg.Page - 1) * pg.Limit
}
