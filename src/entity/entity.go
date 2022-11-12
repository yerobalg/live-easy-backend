package entity

type PaginationParam struct {
	Limit int `form:"limit" param:"limit"`
	Page  int `form:"page" param:"page"`
}
