package entity

import (
	"math"
)

type HTTPResponse struct {
	Meta       Meta             `json:"metaData"`
	Message    string           `json:"message"`
	IsSuccess  bool             `json:"isSuccess"`
	Data       interface{}      `json:"data"`
	Pagination *PaginationParam `json:"pagination"`
}

type Meta struct {
	Time        string  `json:"timestamp"`
	RequestID   string `json:"requestId"`
	TimeElapsed string `json:"timeElapsed"`
}

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

func FormatPaginationParam(params PaginationParam) PaginationParam {
	paginationParam := params

	if params.Limit == 0 {
		paginationParam.Limit = 10
	}

	if params.Page == 0 {
		paginationParam.Page = 1
	}

	return paginationParam
}
