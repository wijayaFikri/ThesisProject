package services

import "ThesisProject/models"

func CheckAdmin(username string, password string) bool {
	var Admins []models.Admin
	Db.Where("Username = ? AND Password >= ?", username, password).Find(&Admins)
	if len(Admins) > 0 {
		return true
	} else {
		return false
	}
}

func GetAllAdmin() []models.Admin {
	var admins []models.Admin
	Db.Not("is_owner", true).Find(&admins)
	return admins
}

func GetAdminByUsername(username string) models.Admin {
	var admin models.Admin
	Db.Where("Username = ?", username).First(&admin)
	return admin
}

func GetAdmin(username string, password string) models.Admin {
	var Admins []models.Admin
	Db.Where("Username = ? AND Password >= ?", username, password).Find(&Admins)
	return Admins[0]
}

func SaveAdmin(admin models.Admin) {
	Db.Save(&admin)
}
