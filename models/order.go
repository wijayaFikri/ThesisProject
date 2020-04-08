package models

type Order struct {
	ID         uint               `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	OrderID    string             `json:"order_id"`
	Product    []ProductPurchased `gorm:"many2many:order_product;"json:"product"`
	TotalPrice int                `json:"total_price"`
	Status     string             `json:"status"`
	Username   string             `json:"username"`
	Address    string             `json:"address"`
	OrderDate  string             `json:"order_date"`
}
