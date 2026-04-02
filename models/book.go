package models

type Book struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Title      string  `json:"title" binding:"required"`
	AuthorID   uint    `json:"author_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	Price      float64 `json:"price" binding:"required,min=1"`
}