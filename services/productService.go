package services

import (
	"ThesisProject/models"
	"strings"
)

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

func GetAllProductByFilter(filter string) []models.Product {
	var products []models.Product
	if filter == "latest" {
		Db.Order("ID desc").Find(&products)
	} else if filter == "lowest_price" {
		Db.Order("price asc").Find(&products)
	} else if filter == "highest_price" {
		Db.Order("price desc").Find(&products)
	} else if strings.Contains(filter, ",") {
		filters := strings.Split(filter, ",")
		secondFilter := filters[1]
		if secondFilter == "latest" {
			Db.Where("Category = ?", filters[0]).Order("ID desc").Find(&products)
		} else if secondFilter == "lowest_price" {
			Db.Where("Category = ?", filters[0]).Order("price asc").Find(&products)
		} else if secondFilter == "highest_price" {
			Db.Where("Category = ?", filters[0]).Order("price desc").Find(&products)
		}
	} else {
		Db.Where("Category = ?", filter).Find(&products)
	}
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

func GetStockLessThanTen() []models.Product {
	var Products []models.Product
	Db.Where("Quantity <= ?", 10).Find(&Products)
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
