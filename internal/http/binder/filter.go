package binder

type Time string

const (
	Today     Time = "today"
	Tomorrow  Time = "tomorrow"
	ThisWeek  Time = "this-week"
	NextWeek  Time = "next-week"
	ThisMonth Time = "this-month"
	NextMonth Time = "next-month"
)

type FilterRequest struct {
	Keyword  *string `json:"keyword" query:"keyword"`
	Location *int    `json:"location" query:"location"`
	Topic    *int    `json:"topic" query:"topic"`
	Category *int    `json:"category" query:"category"`
	Time     *Time   `json:"time" query:"time"`
	IsPaid   *bool   `json:"is_paid" query:"is_paid"`
}
