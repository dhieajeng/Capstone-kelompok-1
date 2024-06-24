package binder

type SortRequest struct {
	Sort *string `json:"sort" query:"sort"`
}
