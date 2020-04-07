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
	allProduct := services.GetAllProduct()
	c.JSON(http.StatusOK, gin.H{
		"products": allProduct,
	})
}

func CreateOrder(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var result map[string]interface{}
	json.Unmarshal([]byte(reqBody), &result)
	data := result["productList"].([]interface{})
	address := result["address"]

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
		order.Product = append(order.Product, purchasedProduct)
		totalPrice = totalPrice + purchasedProduct.Price
	}
	order.Status = STATUS_OPEN
	order.TotalPrice = totalPrice
	user := services.FindUserById(1)
	user.Order = append(user.Order, order)
	services.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"address": address,
	})
}