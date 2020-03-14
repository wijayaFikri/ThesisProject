package models

type Product struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	VendorID uint   `json:"vendor_id"`
}
