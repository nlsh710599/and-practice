package model

type Number struct {
	Name  string `json:"name" gorm:"primarykey"`
	Value string `json:"value"`
}
