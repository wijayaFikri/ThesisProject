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
	Db.Preload("Product").Find(Db.Find(&Vendors))
	return Vendors
}

func FindVendorById(id int) models.Vendor {
	Vendor := models.Vendor{ID: id}
	Db.Preload("Product").Find(&Vendor)
	return Vendor
}
