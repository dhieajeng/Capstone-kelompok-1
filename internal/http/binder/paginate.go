package binder

type PaginateRequest struct {
	Page  *int `json:"page" query:"page"`
	Limit *int `json:"limit" query:"limit"`
}
