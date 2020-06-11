package services

import "ThesisProject/models"

func AddVendorPurchase(vendorPurchase models.VendorPurchase) {
	Db.Save(&vendorPurchase)
}

func FindVendorPurchaseById(id int) models.VendorPurchase {
	vendorPurchase := models.VendorPurchase{ID: id}
	Db.Preload("Product").Find(&vendorPurchase)
	return vendorPurchase
}

func GetAllVendorPurchase() []models.Vendor {
	var Vendors []models.Vendor
	Db.Preload("Product").Find(&Vendors)
	return Vendors
}

func FindVendorPurchaseByVendorName(name string) []models.VendorPurchase {
	var vendorPurchase []models.VendorPurchase
	Db.Preload("product").Where("vendor_name = ?", name).Find(&vendorPurchase)
	return vendorPurchase
}
