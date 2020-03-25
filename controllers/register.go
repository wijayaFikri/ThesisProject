package controllers

import (
	"ThesisProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "/register/register-2.html", gin.H{})
}

func Test(c *gin.Context) {
	listProduct := services.GetAllProduct()
	listVendor := services.GetAllVendor()
	listOrder := services.GetAllOrder()
	listUser := services.GetAllUser()
	c.JSON(http.StatusOK, gin.H{
		"ping":   "pong",
		"pdt":    listProduct,
		"order":  listOrder,
		"vendor": listVendor,
		"user":   listUser,
	})
}
