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
