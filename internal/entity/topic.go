package entity

type Topic struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Auditable
}

func NewTopic(name string) *Topic {
	return &Topic{Name: name}
}
