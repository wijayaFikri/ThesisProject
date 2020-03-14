package models

import "time"

type Admin struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT";json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
}
