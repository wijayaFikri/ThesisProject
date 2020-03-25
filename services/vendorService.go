package services

import "ThesisProject/models"

func AddVendor(Vendor models.Vendor) {
	Db.Save(&Vendor)
}

func RemoveVendor(Vendor models.Vendor) {
	Db.Delete(&Vendor)
}

func UpdateVendor(Vendor models.Vendor) {
	Db.Update(&Vendor)
}

func GetAllVendor() []models.Vendor {
	var Vendors []models.Vendor
	Db.Preload("Product").Find(&Vendors)
	return Vendors
}

func GetLatestVendor() []models.Vendor {
	var Vendors []models.Vendor
	Db.Preload("Product").Order("ID desc").Limit(7).Find(&Vendors)
	return Vendors
}

func GetAllVendorName() []string {
	var Vendors []models.Vendor
	Db.Select("vendor_name").Find(&Vendors)
	var vendorName []string
	for i := 0; i < len(Vendors); i++ {
		vendorName = append(vendorName, Vendors[i].VendorName)
	}
	return vendorName
}

func FindVendorById(id int) models.Vendor {
	Vendor := models.Vendor{ID: id}
	Db.Preload("Product").Find(&Vendor)
	return Vendor
}

func FindVendorByName(name string) models.Vendor {
	vendor := models.Vendor{VendorName: name}
	Db.Preload("product").Where("vendor_name = ?", name).Find(&vendor)
	return vendor
}
