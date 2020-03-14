package services

import "ThesisProject/models"

func AddOrder(Order models.Order) {
	Db.Save(&Order)
}

func RemoveOrder(Order models.Order) {
	Db.Delete(&Order)
}

func UpdateOrder(Order models.Order) {
	Db.Update(&Order)
}

func GetAllOrder() []models.Order {
	var Orders []models.Order
	Db.Preload("Product").Find(Db.Find(&Orders))
	return Orders
}

func FindOrderById(id uint) models.Order {
	Order := models.Order{ID: id}
	Db.Preload("Product").Find(&Order)
	return Order
}
