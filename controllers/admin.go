package controllers

import (
	"ThesisProject/models"
	"ThesisProject/services"
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

const USERNAME = "username"
const VENDOR = "vendor"
const PRODUCT = "product"
const USER = "user"
const ORDER = "order"

func AdminLogin(c *gin.Context) {
	isError := false
	errorMessage := ""
	if c.PostForm("username") != "" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		result := services.CheckAdmin(username, password)
		if result {
			session := sessions.Default(c)
			session.Set(USERNAME, username)
			session.Save()
			c.Redirect(http.StatusFound, "/admin/dashboard")
			return
		} else {
			errorMessage = "Wrong Username or Password"
			isError = true
		}
	}
	c.HTML(http.StatusOK, "/admin/login.html", gin.H{
		"isError":      isError,
		"errorMessage": errorMessage,
	})
}

func AdminDashboard(c *gin.Context) {
	if !checkLogin(c) {
		c.Redirect(http.StatusFound, "/admin")
		return
	}
	latestOrder := services.GetLatestOrder()
	latestVendor := services.GetLatestVendor()
	latestProduct := services.GetLatestProduct()
	c.HTML(http.StatusOK, "/admin/dashboard.html", gin.H{
		"latestProduct": latestProduct,
		"latestVendor":  latestVendor,
		"latestOrder":   latestOrder,
		"activeMenu":    "",
	})
}

func checkLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	user := session.Get(USERNAME)
	if user == nil {
		return false
	} else {
		return true
	}
}

func InventoryList(c *gin.Context) {
	status := checkLogin(c)
	if !status {
		c.Redirect(http.StatusFound, "/admin")
		return
	}
	resultMessage := ""
	resultError := ""
	result := false
	isError := false
	isDeleted := false
	resultDelete := ""
	searchMessage := ""
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
		file, header, _ := c.Request.FormFile("image")
		filename := header.Filename
		product.ImageUrl = uploadToImgur(file, filename)
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
	var allProduct []models.Product
	if c.PostForm("searchKey") != "" {
		allProduct = services.SearchProduct(c.PostForm("searchKey"))
		if !(len(allProduct) > 0) {
			searchMessage = "Product not found"
			isError = true
		}
	} else {
		allProduct = services.GetAllProduct()
	}

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
		"activeMenu":           PRODUCT,
		"searchMessage":        searchMessage,
	})
}

func uploadToImgur(image io.Reader, filename string) string {
	var buf = new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	part, _ := writer.CreateFormFile("image", filename)
	io.Copy(part, image)

	writer.Close()
	req, _ := http.NewRequest("POST", "https://api.imgur.com/3/image", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Client-ID "+"fca6e3aed6b3058")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)
	data := result["data"].(map[string]interface{})
	return data["link"].(string)
}

func ShowVendors(c *gin.Context) {
	if !checkLogin(c) {
		c.Redirect(http.StatusFound, "/admin")
		return
	}
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
		"activeMenu":    VENDOR,
	})
}

func EditDetailProduct(c *gin.Context) {
	if !checkLogin(c) {
		c.Redirect(http.StatusFound, "/admin")
		return
	}
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
		"activeMenu":    PRODUCT,
	})
}

func ShowDetailVendor(c *gin.Context) {
	if !checkLogin(c) {
		c.Redirect(http.StatusFound, "/admin")
		return
	}
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
			"activeMenu":    VENDOR,
		})
	}
}

func ShowUsers(c *gin.Context) {
	isError := false
	resultMessage := ""
	if c.PostForm("username") != "" {
		var user models.User
		user.Username = c.PostForm("username")
		user.Password = c.PostForm("password")
		user.FirstName = c.PostForm("firstName")
		user.LastName = c.PostForm("lastName")
		user.Address = c.PostForm("address")
		if services.CheckUser(user) {
			isError = true
			resultMessage = "User already exist"
		} else {
			services.AddUser(user)
			resultMessage = "New User added successfully"
		}
	}

	if c.PostForm("id") != "" {
		id, err := strconv.Atoi(c.PostForm("id"))
		if err == nil {
			product := services.FindUserById(uint(id))
			services.RemoveUser(product)
			resultMessage = "Data deleted successfully"
		}

	}

	var users []models.User
	if c.PostForm("searchKey") != "" {
		users = services.SearchUser(c.PostForm("searchKey"))
		if !(len(users) > 0) {
			resultMessage = "User not found"
			isError = true
		}
	} else {
		users = services.GetAllUser()
	}
	c.HTML(http.StatusOK, "admin/users.html", gin.H{
		"activeMenu":    USER,
		"allUser":       users,
		"isError":       isError,
		"resultMessage": resultMessage,
	})
}

func ShowOrder(c *gin.Context) {
	allOrder := services.GetAllOrder()
	c.HTML(http.StatusOK, "/admin/orders.html", gin.H{
		"allOrder":   allOrder,
		"activeMenu": ORDER,
	})
}

func ShowOrderDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	order := services.FindOrderById(uint(id))
	productSize := len(order.Product)
	c.HTML(http.StatusOK, "/admin/order_detail.html", gin.H{
		"order":       order,
		"activeMenu":  ORDER,
		"productSize": productSize,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/admin")
}
