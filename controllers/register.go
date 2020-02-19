package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "/register/register.html", gin.H{})
}
