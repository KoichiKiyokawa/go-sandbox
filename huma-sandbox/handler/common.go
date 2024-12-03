package handler

type PaginationInput struct {
	Per  int `query:"per"`
	Page int `query:"page"`
}
