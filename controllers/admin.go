package controllers

import (
	"ThesisProject/models"
	"ThesisProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func AdminLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "/admin/login.html", gin.H{})
}

func AdminDashboard(c *gin.Context) {
	latestOrder := services.GetLatestOrder()
	latestVendor := services.GetLatestVendor()
	latestProduct := services.GetLatestProduct()
	c.HTML(http.StatusOK, "/admin/dashboard.html", gin.H{
		"user":          "ngentot",
		"latestProduct": latestProduct,
		"latestVendor":  latestVendor,
		"latestOrder":   latestOrder,
	})
}

func InventoryList(c *gin.Context) {
	resultMessage := ""
	resultError := ""
	result := false
	isError := false
	isDeleted := false
	resultDelete := ""
	if c.PostForm("id") != "" {
		id, err := strconv.Atoi(c.PostForm("id"))
		if err == nil {
			product := services.FindProductById(uint(id))
			services.RemoveProduct(product)
			resultDelete = "Data deleted successfully"
			isDeleted = true
		}

	}
	if c.PostForm("name") != "" {
		var product models.Product
		var convertError error
		product.Name = c.PostForm("name")
		product.Quantity, convertError = strconv.Atoi(c.PostForm("quantity"))
		product.Price, convertError = strconv.Atoi(c.PostForm("price"))
		vendor := services.FindVendorByName(c.PostForm("vendorName"))
		vendor.Product = append(vendor.Product, product)
		if convertError == nil {
			services.Db.Save(&vendor)
		} else {
			isError = true
			resultError = convertError.Error()
		}
		resultMessage = "Data added successfully"
		result = true
	}
	allProduct := services.GetAllProduct()
	mostPurchasedProduct := services.GetMostPurchasedProduct()
	vendorNames := services.GetAllVendorName()
	c.HTML(http.StatusOK, "admin/inventories.html", gin.H{
		"allProduct":           allProduct,
		"mostPurchasedProduct": mostPurchasedProduct,
		"vendorNames":          vendorNames,
		"resultMessage":        resultMessage,
		"resultError":          resultError,
		"isResult":             result,
		"isError":              isError,
		"isDeleted":            isDeleted,
		"resultDelete":         resultDelete,
	})
}

func ShowVendors(c *gin.Context) {
	resultMessage := ""
	resultError := ""
	result := false
	isError := false
	if c.PostForm("name") != "" {
		vendor := models.Vendor{VendorName: c.PostForm("name"), VendorLocation: c.PostForm("location"),
			VendorDescription: c.PostForm("description")}
		services.AddVendor(vendor)
		resultMessage = "Data added successfully"
		result = true
		isError = false
	}
	latestVendor := services.GetLatestVendor()
	allVendor := services.GetAllVendor()
	c.HTML(http.StatusOK, "admin/vendors.html", gin.H{
		"latestVendor":  latestVendor,
		"allVendor":     allVendor,
		"resultMessage": resultMessage,
		"resultError":   resultError,
		"isResult":      result,
		"isError":       isError,
	})
}

func EditDetailProduct(c *gin.Context) {
	isResult := false
	resultMessage := ""
	id, err := strconv.Atoi(c.PostForm("id"))
	var product models.Product
	if err == nil {
		product = services.FindProductById(uint(id))
		if c.PostForm("name") != "" {
			product.Name = c.PostForm("name")
			product.Quantity, _ = strconv.Atoi(strings.TrimSpace(c.PostForm("quantity")))
			product.Price, _ = strconv.Atoi(strings.TrimSpace(c.PostForm("price")))
			product.VendorName = c.PostForm("vendorName")
			services.Db.Save(&product)
			isResult = true
			resultMessage = "Data Changed Successfully"
		}
	}
	vendorNames := services.GetAllVendorName()
	c.HTML(http.StatusOK, "admin/product_detail.html", gin.H{
		"product":       product,
		"vendorNames":   vendorNames,
		"isResult":      isResult,
		"resultMessage": resultMessage,
	})
}

func ShowDetailVendor(c *gin.Context) {
	isResult := false
	resultMessage := ""
	id, err := strconv.Atoi(c.PostForm("id"))
	if c.PostForm("name") != "" {
		updatedVendor := services.FindVendorById(id)
		updatedVendor.VendorName = c.PostForm("name")
		updatedVendor.VendorLocation = c.PostForm("description")
		updatedVendor.VendorDescription = c.PostForm("location")
		services.Db.Save(&updatedVendor)
		isResult = true
		resultMessage = "Changes saved successfully"
	}
	vendor := services.FindVendorById(id)
	if err == nil {
		c.HTML(http.StatusOK, "admin/vendor_detail.html", gin.H{
			"vendor":        vendor,
			"isResult":      isResult,
			"resultMessage": resultMessage,
		})
	}
}
