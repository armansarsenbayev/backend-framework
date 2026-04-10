package models

type Restaurant struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
