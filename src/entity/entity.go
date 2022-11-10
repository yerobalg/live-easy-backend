package entity

type PaginationParam struct {
	Limit  int `form:"limit" param:"limit"`
	Offset int `form:"page" param:"page"`
}
