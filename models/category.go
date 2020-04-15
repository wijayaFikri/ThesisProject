package models

type Category struct {
	ID           int     `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	CategoryName string  `json:"category_name"`
	Product      Product `gorm:"foreignkey:Category;association_foreignkey:CategoryName";json:"product"`
}
