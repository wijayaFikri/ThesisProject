package models

type Order struct {
	ID         uint      `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	OrderID    string    `json:"order_id"`
	Product    []Product `gorm:"many2many:order_product;"`
	TotalPrice int       `json:"total_price"`
}
