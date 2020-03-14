package services

import "ThesisProject/models"

func AddUser(User models.User) {
	Db.Save(&User)
}

func RemoveUser(User models.User) {
	Db.Delete(&User)
}

func UpdateUser(User models.User) {
	Db.Update(&User)
}

func GetAllUser() []models.User {
	var Users []models.User
	Db.Find(Db.Find(&Users))
	return Users
}

func FindUserById(id uint) models.User {
	User := models.User{ID: id}
	Db.Find(&User)
	return User
}
