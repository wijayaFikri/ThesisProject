package controllers

import (
	"ThesisProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendAllProduct(c *gin.Context) {
	allProduct := services.GetAllProduct()
	c.JSON(http.StatusOK, gin.H{
		"products": allProduct,
	})
}
