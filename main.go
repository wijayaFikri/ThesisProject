package main

import (
	"ThesisProject/controllers"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/", controllers.Index)
	r.Static("/assets","./assets")
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
