package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null" binding:"required"`
	Password string `json:"password" binding:"required"`
}
