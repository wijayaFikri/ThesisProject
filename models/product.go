package models

import (
	"time"
)

type Product struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	Name       string `json:"name";form:"name"`
	Price      int    `json:"price";form:"price"`
	Quantity   int    `json:"quantity";form:"quantity"`
	VendorName string `json:"vendor_name";form:"vendorName"`
	ImageUrl   string `json:"image_url"`
	Purchased  int    `json:"purchased";form:"purchased"`
}

type ProductPurchased struct {
	ID           uint      `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	ProductId    uint      `json:"product_id"`
	Name         string    `json:"name"`
	Price        int       `json:"price"`
	Quantity     int       `json:"quantity"`
	VendorName   string    `json:"vendor_name"`
	ImageUrl     string    `json:"image_url"`
	PurchaseDate time.Time `json:"purchase_date"`
}
