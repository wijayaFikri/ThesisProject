package services

import "ThesisProject/models"

func AddOrder(Order models.Order) {
	Db.Save(&Order)
}

func RemoveOrder(Order models.Order) {
	Db.Delete(&Order)
}

func UpdateOrder(Order models.Order) {
	Db.Save(&Order)
}

func GetAllOrder() []models.Order {
	var Orders []models.Order
	Db.Preload("Product").Find(&Orders)
	return Orders
}

func GetAllOrderSortByDate(sortKey string) []models.Order {
	var Orders []models.Order
	Db.Order("Created_Date " + sortKey).Find(&Orders)
	return Orders
}

func GetLatestOrder() []models.Order {
	var Orders []models.Order
	Db.Preload("Product").Order("ID desc").Limit(7).Find(&Orders)
	return Orders
}

func FindOrderById(id uint) models.Order {
	Order := models.Order{ID: id}
	Db.Preload("Product").Find(&Order)
	return Order
}
