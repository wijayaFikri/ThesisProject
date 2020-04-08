package services

import "ThesisProject/models"

func AddUser(User models.User) {
	Db.Save(&User)
}

func RemoveUser(User models.User) {
	Db.Delete(&User)
}

func UpdateUser(User models.User) {
	Db.Save(&User)
}

func GetAllUser() []models.User {
	var Users []models.User
	Db.Preload("Order").Find(&Users)
	return Users
}

func FindUserById(id uint) models.User {
	User := models.User{ID: id}
	Db.Preload("Order.Product").Find(&User)
	return User
}

func GetLatestUser() []models.User {
	var users []models.User
	Db.Order("ID desc").Limit(7).Find(&users)
	return users
}

func SearchUser(searchKey string) []models.User {
	var users []models.User
	key := "%" + searchKey + "%"
	Db.Where("First_Name LIKE ? OR Last_Name LIKE ?", key, key).Find(&users)
	return users
}

func LoginUser(username string, password string) (models.User, bool) {
	var users []models.User
	Db.Where("Username = ? AND Password >= ?", username, password).Find(&users)
	if len(users) > 0 {
		return users[0], true
	} else {
		return models.User{}, false
	}
}

func CheckUser(user models.User) bool {
	var users []models.User
	Db.Where("Username = ? AND Password >= ?", user.Username, user.Password).Find(&users)
	if len(users) > 0 {
		return true
	} else {
		return false
	}
}

func GetMostActiveUser() []models.User {
	var users []models.User
	Db.Order("TotalOrder desc").Limit(8).Find(&users)
	return users
}
