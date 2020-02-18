package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadLoginScreen(c *gin.Context) {
	c.HTML(http.StatusOK, "/authentication/login_screen.html", gin.H{})
}
