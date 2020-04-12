package models

type Admin struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsOwner   bool   `json:"is_owner"`
}
