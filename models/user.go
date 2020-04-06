package models

type User struct {
	ID        uint    `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Location  string  `json:"location"`
	Order     []Order `gorm:"foreignkey:Username;association_foreignkey:Username";json:"order";form:"product"`
}
