package main

import (
	"ThesisProject/controllers"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	ginView := ginview.New(goview.Config{
		Root: "views",
		Funcs: template.FuncMap{
			"add": func(x, y int) int {
				return x + y
			},
			"equal": func(x, y string) bool {
				if x == y {
					return true
				} else {
					return false
				}
			},
			"isExist": func(message string) bool {
				if message != "" {
					return true
				} else {
					return false
				}
			},
		},
	})
	r.Use(sessions.Sessions("mySession", store))
	r.HTMLRender = ginView
	r.GET("/", controllers.AdminDashboard)
	r.POST("/admin", controllers.AdminLogin)
	r.GET("/login", controllers.LoadLoginScreen)
	r.GET("/register", controllers.Register)
	r.GET("/admin", controllers.AdminLogin)
	r.GET("/admin/dashboard", controllers.AdminDashboard)
	r.POST("/admin/dashboard", controllers.AdminDashboard)
	r.GET("/admin/inventory", controllers.InventoryList)
	r.POST("/admin/inventory", controllers.InventoryList)
	r.GET("admin/vendor", controllers.ShowVendors)
	r.POST("/admin/vendor", controllers.ShowVendors)
	r.POST("/admin/vendor/detail", controllers.ShowDetailVendor)
	r.POST("/admin/inventory/detail", controllers.EditDetailProduct)
	r.GET("/admin/user", controllers.ShowUsers)
	r.POST("/admin/user", controllers.ShowUsers)
	r.GET("/admin/order", controllers.ShowOrder)
	r.POST("/admin/order/detail", controllers.ShowOrderDetail)
	r.POST("/admin/order/result", controllers.ProcessOrder)
	r.GET("/logout", controllers.Logout)
	r.GET("/test", controllers.Test)
	r.GET("/admin/employee/manage", controllers.EmployeeManagement)
	r.POST("/admin/employee/manage", controllers.EmployeeManagement)
	r.POST("/vendor/detail/add/product", controllers.ShowAddMoreQuantityVendor)

	//API BUAT MOBILE
	r.GET("/mobile/products", controllers.SendAllProduct)
	r.POST("/mobile/create/order", controllers.CreateOrder)
	r.POST("/mobile/user/login", controllers.LoginUser)
	r.POST("mobile/order/history", controllers.SendOrderHistory)
	r.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run(":8080")
	}
}
