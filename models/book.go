package models

// Book struct model
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key`
	Title  string `json:"title"`
	Author string `json:"author"`
}
