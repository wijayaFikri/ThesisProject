package main

import (
	"ThesisProject/controllers"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mySession", store))
	r.HTMLRender = ginview.Default()
	r.GET("/", controllers.Index)
	r.GET("/login", controllers.LoadLoginScreen)
	r.GET("/register", controllers.Register)
	r.GET("/admin", controllers.AdminLogin)
	r.GET("/admin/dashboard", controllers.AdminDashboard)
	r.GET("/test", controllers.Test)
	r.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run(":8080")
	}
}
