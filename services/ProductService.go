package services

import "ThesisProject/models"

func AddProduct(product models.Product) {
	Db.Save(&product)
}

func RemoveProduct(product models.Product) {
	Db.Delete(&product)
}

func UpdateProduct(product models.Product) {
	Db.Update(&product)
}

func GetAllProduct() []models.Product {
	var products []models.Product
	Db.Find(Db.Find(&products))
	return products
}

func FindProductById(id uint) models.Product {
	product := models.Product{ID: id}
	Db.Find(&product)
	return product
}
