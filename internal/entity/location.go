package entity

type Location struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Auditable
}

func NewLocation(name string) *Location {
	return &Location{Name: name}
}
