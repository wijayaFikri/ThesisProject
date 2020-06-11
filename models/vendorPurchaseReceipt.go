package models

type VendorPurchase struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	OrderID    string    `json:"order_id"`
	Date       string    `json:"date"`
	VendorName string    `json:"vendor_name"`
	Product    []Product `json:"product"`
}

type VendorPurchaseDisplay struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
