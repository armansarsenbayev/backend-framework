package models

type Order struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	RestaurantID uint    `json:"restaurant_id" binding:"required"`
	CustomerName string  `json:"customer_name" binding:"required"`
	Status       string  `json:"status"`
	TotalPrice   float64 `json:"total_price" binding:"required,min=1"`
}
