package models

type User struct {
	ID         uint    `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Address    string  `json:"address"`
	Order      []Order `gorm:"foreignkey:Username;association_foreignkey:Username";json:"order";form:"product"`
	TotalOrder int     `json:"total_order"`
}

type UserForm struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Changes   string `json:"changes"`
}

func ConvertFormToUser(user User, form UserForm) User {
	user.ID = form.ID
	user.Password = form.Password
	user.Username = form.Username
	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.Address = form.Address
	return user
}
