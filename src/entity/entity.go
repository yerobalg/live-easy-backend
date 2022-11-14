package entity

type PaginationParam struct {
	Limit        int `form:"limit" param:"limit" json:"-"`
	Page         int `form:"page" param:"page" json:"-"`
	CurrentPage  int `param:"current_page" json:"currentPage"`
	TotalPage    int `param:"total_page" json:"totalPage"`
	TotalElement int `param:"total_element" json:"totalElement"`
}
