package entity

import (
	"math"
)

type PaginationParam struct {
	Limit          int64 `form:"limit" json:"limit" gorm:"-"`
	Page           int64 `form:"page" json:"-" gorm:"-"`
	Offset         int64 `json:"-" gorm:"-"`
	CurrentPage    int64 `json:"currentPage" gorm:"-"`
	TotalPage      int64 `json:"totalPage" gorm:"-"`
	CurrentElement int64 `json:"currentElement" gorm:"-"`
	TotalElement   int64 `json:"totalElement" gorm:"-"`
}

func (pg *PaginationParam) ProcessPagination(rowsAffected int64) {
	pg.CurrentPage = pg.Page
	pg.TotalPage = int64(math.Ceil(float64(pg.TotalElement) / float64(pg.Limit)))
	pg.CurrentElement = rowsAffected
}

func (pg *PaginationParam) SetLimitOffset() {
	if pg.Limit < 1 {
		pg.Limit = 10
	}

	if pg.Page < 1 {
		pg.Page = 1
	}

	pg.Offset = (pg.Page - 1) * pg.Limit
}
