package services

import "ThesisProject/models"

func AddProduct(product models.Product) {
	Db.Save(&product)
}

func RemoveProduct(product models.Product) {
	Db.Delete(&product)
}

func UpdateProduct(product models.Product) {
	Db.Save(&product)
}

func GetAllProduct() []models.Product {
	var products []models.Product
	Db.Find(&products)
	return products
}

func SearchProduct(key string) []models.Product {
	var products []models.Product
	key = "%" + key + "%"
	Db.Where("Name LIKE ?", key).Find(&products)
	return products
}

func GetLatestProduct() []models.Product {
	var Products []models.Product
	Db.Order("ID desc").Limit(7).Find(&Products)
	return Products
}

func GetMostPurchasedProduct() []models.Product {
	var Products []models.Product
	Db.Order("Purchased desc").Limit(8).Find(&Products)
	return Products
}

func FindProductById(id uint) models.Product {
	product := models.Product{ID: id}
	Db.Find(&product)
	return product
}
