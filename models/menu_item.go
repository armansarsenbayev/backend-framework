package models

type MenuItem struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	RestaurantID uint    `json:"restaurant_id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Price        float64 `json:"price" binding:"required,min=1"`
}
