package models

type Vendor struct {
	ID                int       `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	VendorName        string    `json:"vendor_name"`
	VendorLocation    string    `json:"vendor_location"`
	VendorDescription string    `json:"vendor_description"`
	Product           []Product `json:"product"`
}
