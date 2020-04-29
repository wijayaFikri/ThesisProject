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
const IS_OWNER = "is_owner"
const VENDOR_ID_KEY = "vendor_id"

func AdminLogin(c *gin.Context) {
	isError := false
	errorMessage := ""
	if c.PostForm("username") != "" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		result := services.CheckAdmin(username, password)
		if result {
			admin := services.GetAdmin(username, password)
			session := sessions.Default(c)
			session.Set(USERNAME, username)
			session.Set(IS_OWNER, admin.IsOwner)
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
	latestUser := services.GetLatestUser()
	stockLessThanTen := services.GetStockLessThanTen()
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "/admin/dashboard.html", gin.H{
		"latestProduct":    latestProduct,
		"latestVendor":     latestVendor,
		"latestOrder":      latestOrder,
		"latestUser":       latestUser,
		"stockLessThanTen": stockLessThanTen,
		"activeMenu":       "",
		"isOwner":          isOwner,
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

func getIsOwner(c *gin.Context) bool {
	session := sessions.Default(c)
	return session.Get(IS_OWNER).(bool)
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
	searchMessage := ""
	if c.PostForm("id") != "" {
		id, err := strconv.Atoi(c.PostForm("id"))
		if err == nil {
			product := services.FindProductById(uint(id))
			services.RemoveProduct(product)
			resultMessage = "Data deleted successfully"
		}

	}

	if c.PostForm("categoryName") != "" {
		var category models.Category
		category.CategoryName = c.PostForm("categoryName")
		services.AddCategory(category)
		resultMessage = "New Category Added Successfully"
	}
	if c.PostForm("name") != "" {
		var product models.Product
		var convertError error
		product.Name = c.PostForm("name")
		product.Category = c.PostForm("category")
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
	} else if c.PostForm("Category") != "" {
		allProduct = services.GetAllProductByFilter(c.PostForm("Category"))
	} else {
		allProduct = services.GetAllProduct()
	}

	mostPurchasedProduct := services.GetMostPurchasedProduct()
	vendorNames := services.GetAllVendorName()
	categoryNames := services.GetAllCategoryName()
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "admin/inventories.html", gin.H{
		"allProduct":           allProduct,
		"mostPurchasedProduct": mostPurchasedProduct,
		"vendorNames":          vendorNames,
		"resultMessage":        resultMessage,
		"resultError":          resultError,
		"isResult":             result,
		"isError":              isError,
		"activeMenu":           PRODUCT,
		"searchMessage":        searchMessage,
		"isOwner":              isOwner,
		"categoryNames":        categoryNames,
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
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "admin/vendors.html", gin.H{
		"latestVendor":  latestVendor,
		"allVendor":     allVendor,
		"resultMessage": resultMessage,
		"resultError":   resultError,
		"isResult":      result,
		"isError":       isError,
		"activeMenu":    VENDOR,
		"isOwner":       isOwner,
	})
}

func EditDetailProduct(c *gin.Context) {
	if !checkLogin(c) {
		c.Redirect(http.StatusFound, "/admin")
		return
	}
	resultMessage := ""
	id, err := strconv.Atoi(c.PostForm("id"))
	var product models.Product
	if err == nil {
		product = services.FindProductById(uint(id))
		if c.PostForm("name") != "" {
			product.Name = c.PostForm("name")
			product.Category = c.PostForm("category")
			product.Price, _ = strconv.Atoi(strings.TrimSpace(c.PostForm("price")))
			product.VendorName = c.PostForm("vendorName")
			services.Db.Save(&product)
			resultMessage = "Data Changed Successfully"
		}
	}
	categoryNames := services.GetAllCategoryName()
	vendorNames := services.GetAllVendorName()
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "admin/product_detail.html", gin.H{
		"product":       product,
		"vendorNames":   vendorNames,
		"resultMessage": resultMessage,
		"activeMenu":    PRODUCT,
		"isOwner":       isOwner,
		"categoryNames": categoryNames,
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
	session := sessions.Default(c)
	session.Set(VENDOR_ID_KEY, id)
	session.Save()
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
	isOwner := getIsOwner(c)
	if err == nil {
		c.HTML(http.StatusOK, "admin/vendor_detail.html", gin.H{
			"vendor":        vendor,
			"isResult":      isResult,
			"resultMessage": resultMessage,
			"activeMenu":    VENDOR,
			"isOwner":       isOwner,
		})
	}
}

func ShowAddMoreQuantityVendor(c *gin.Context) {
	if !checkLogin(c) {
		c.Redirect(http.StatusFound, "/admin")
		return
	}
	resultMessage := ""
	if c.PostForm("payload") != "" {
		payload := c.PostForm("payload")
		var result []map[string]string
		json.Unmarshal([]byte(payload), &result)
		for _, element := range result {
			id, _ := strconv.Atoi(element["id"])
			product := services.FindProductById(uint(id))
			addedQuantity, _ := strconv.Atoi(element["purchase_quantity"])
			product.Quantity = product.Quantity + addedQuantity
			services.UpdateProduct(product)
		}

		resultMessage = "Product purchased Successfully"
	}
	session := sessions.Default(c)

	id := session.Get(VENDOR_ID_KEY).(int)
	vendor := services.FindVendorById(id)
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "admin/vendor_add_more_quantity.html", gin.H{
		"resultMessage": resultMessage,
		"activeMenu":    VENDOR,
		"isOwner":       isOwner,
		"vendor":        vendor,
	})
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
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "admin/users.html", gin.H{
		"activeMenu":    USER,
		"allUser":       users,
		"isError":       isError,
		"resultMessage": resultMessage,
		"isOwner":       isOwner,
	})
}

func ShowOrder(c *gin.Context) {
	desc := true
	var allOrder []models.Order
	if c.PostForm("sortKey") != "" {
		if c.PostForm("sortKey") == "desc" {
			desc = true
		} else {
			desc = false
		}
		allOrder = services.GetAllOrderSortByDate(c.PostForm("sortKey"))
	} else {
		allOrder = services.GetAllOrder()
	}
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "/admin/orders.html", gin.H{
		"allOrder":   allOrder,
		"activeMenu": ORDER,
		"isOwner":    isOwner,
		"desc":       desc,
	})
}

func ShowOrderDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	order := services.FindOrderById(uint(id))
	productSize := len(order.Product)
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "/admin/order_detail.html", gin.H{
		"order":       order,
		"activeMenu":  ORDER,
		"productSize": productSize,
		"isOwner":     isOwner,
	})
}

func ProcessOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	order := services.FindOrderById(uint(id))
	if order.Status == STATUS_OPEN {
		order.Status = STATUS_PROCESSED
	} else if order.Status == STATUS_PROCESSED {
		order.Status = STATUS_SENT
	} else if order.Status == STATUS_SENT {
		order.Status = STATUS_RECEIVED
	}
	services.UpdateOrder(order)
	resultMessage := "The order status of " + order.OrderID + " has been changed to " + order.Status + " Successfully"
	productSize := len(order.Product)
	isOwner := getIsOwner(c)
	c.HTML(http.StatusOK, "/admin/order_result.html", gin.H{
		"order":         order,
		"activeMenu":    ORDER,
		"productSize":   productSize,
		"resultMessage": resultMessage,
		"isOwner":       isOwner,
	})
}

func EmployeeManagement(c *gin.Context) {
	resultMessage := ""
	session := sessions.Default(c)
	usernameSession := session.Get(USERNAME).(string)
	admin := services.GetAdminByUsername(usernameSession)
	if !checkLogin(c) {
		c.Redirect(http.StatusFound, "/admin")
		return
	}

	if c.PostForm("ownerUsername") != "" {
		newOwnerUsername := c.PostForm("ownerUsername")
		newOwnerPassword := c.PostForm("ownerPassword")
		admin.Username = newOwnerUsername
		admin.Password = newOwnerPassword
		services.SaveAdmin(admin)
		resultMessage = "Changes saved Successfully"
	}

	if c.PostForm("firstname") != "" {
		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
		username := c.PostForm("username")
		password := c.PostForm("password")
		newAdmin := models.Admin{FirstName: firstname, LastName: lastname, Username: username, Password: password}
		services.SaveAdmin(newAdmin)
		resultMessage = "New Employee added Successfully"
	}

	isOwner := getIsOwner(c)
	admins := services.GetAllAdmin()
	c.HTML(http.StatusOK, "/admin/employee_management.html", gin.H{
		"isOwner":       isOwner,
		"activeMenu":    "",
		"allAdmin":      admins,
		"resultMessage": resultMessage,
		"admin":         admin,
	})
}

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("ID"))
	user := services.FindUserById(uint(id))
	isOwner := getIsOwner(c)
	resultMessage := ""

	if c.PostForm("Changes") != "" {
		var newUser models.UserForm
		c.Bind(&newUser)
		user.Password = newUser.Password
		user.Username = newUser.Username
		user.FirstName = newUser.FirstName
		user.LastName = newUser.LastName
		user.Address = newUser.Address
		services.AddUser(user)
		resultMessage = "Changes saved Successfully"
	}
	c.HTML(http.StatusOK, "/admin/user_detail_edit.html", gin.H{
		"isOwner":       isOwner,
		"activeMenu":    USER,
		"resultMessage": resultMessage,
		"user":          user,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/admin")
}
