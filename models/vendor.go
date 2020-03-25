package models

type Vendor struct {
	ID                int       `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	VendorName        string    `json:"vendor_name";form:"vendorName";gorm:"unique;not null"`
	VendorLocation    string    `json:"vendor_location";form:"vendorLocation"`
	VendorDescription string    `json:"vendor_description";form:"vendorDescription"`
	Product           []Product `gorm:"foreignkey:VendorName;association_foreignkey:VendorName";json:"product";form:"product"`
}
