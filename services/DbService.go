package services

import (
	"ThesisProject/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	//open a db connection
	var err error
	Db, err = gorm.Open("mysql", "root:password@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	Db.AutoMigrate(&models.Admin{},
		&models.Product{}, &models.User{}, &models.Vendor{}, &models.Order{})
}
