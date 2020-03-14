package controllers

import (
	"ThesisProject/models"
	"ThesisProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "/register/register-2.html", gin.H{})
}

func Test(c *gin.Context) {
	var vendor models.Vendor
	var product models.Product
	var listVendor []models.Product
	listData := services.Db.Model(&vendor).Find(&vendor)
	listData2 := services.Db.Preload("Vendor").Find(&product)
	services.Db.Find(&listVendor)
	print(listData)
	order := models.Order{OrderID: "515QWE", Product: listVendor, TotalPrice: 500}
	services.Db.Save(&order)
	order2 := services.Db.Preload("Product").First(&models.Order{})
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
		"data": listData,
		"2":    listData2,
		"3":    order2,
	})
}
