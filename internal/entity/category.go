package entity

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Auditable
}

func NewCategory(name string) *Category {
	return &Category{Name: name}
}
