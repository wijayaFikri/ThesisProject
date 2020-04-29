package controllers

import (
	"ThesisProject/models"
	"ThesisProject/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const STATUS_OPEN = "Open"
const STATUS_PROCESSED = "Processed"
const STATUS_SENT = "Sent"
const STATUS_RECEIVED = "Received"

func SendAllProduct(c *gin.Context) {
	var allProduct []models.Product
	if c.PostForm("filter") != "" {
		allProduct = services.GetAllProductByFilter(c.PostForm("filter"))
	} else {
		allProduct = services.GetAllProduct()
	}

	c.JSON(http.StatusOK, gin.H{
		"products": allProduct,
	})
}

func LoginUser(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var result map[string]interface{}
	json.Unmarshal([]byte(reqBody), &result)
	username := result["username"].(string)
	password := result["password"].(string)
	user, _ := services.LoginUser(username, password)
	c.JSON(http.StatusOK, gin.H{
		USER: user,
	})
}

func CreateOrder(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var result map[string]interface{}
	json.Unmarshal([]byte(reqBody), &result)
	data := result["product"].([]interface{})
	address := result["Address"].(string)

	var order models.Order
	orderID := ""
	for i := 0; i < 8; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		random := rand.Intn(10) + (i * 2)
		if random > 10 {
			random = random - 10
		}
		orderID = orderID + strconv.Itoa(random)
	}
	order.OrderID = orderID
	totalPrice := 0
	for _, element := range data {
		productInterface := element.(map[string]interface{})
		var purchasedProduct models.ProductPurchased
		id := productInterface["id"].(float64)
		purchasedProduct.ProductId = uint(id)
		purchasedProduct.VendorName = productInterface["vendor_name"].(string)
		purchasedProduct.Quantity, _ = strconv.Atoi(productInterface["orderQuantity"].(string))
		purchasedProduct.Price = int(productInterface["price"].(float64))
		purchasedProduct.ImageUrl = productInterface["image_url"].(string)
		purchasedProduct.Name = productInterface["name"].(string)
		purchasedProduct.PurchaseDate = time.Now()
		product := services.FindProductById(uint(id))
		product.Quantity = product.Quantity - purchasedProduct.Quantity
		product.Purchased = product.Purchased + purchasedProduct.Quantity
		services.UpdateProduct(product)
		order.Product = append(order.Product, purchasedProduct)
		totalPrice = totalPrice + (purchasedProduct.Price * purchasedProduct.Quantity)
	}
	order.Status = STATUS_OPEN
	order.TotalPrice = totalPrice
	order.OrderDate = time.Now().Format("2006-January-02")
	order.CreatedDate = time.Now()
	order.Address = address
	userId := uint(result["userId"].(float64))
	user := services.FindUserById(userId)
	user.Order = append(user.Order, order)
	user.TotalOrder = user.TotalOrder + 1
	services.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"address": address,
	})
}

func SendOrderHistory(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	user := services.FindUserById(uint(userId))
	c.JSON(http.StatusOK, gin.H{
		"orders": user.Order,
	})
}

func SendCategories(c *gin.Context) {
	categories := services.GetAllCategoryName()
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
