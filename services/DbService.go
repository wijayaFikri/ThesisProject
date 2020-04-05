package services

import (
	"ThesisProject/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var Db *gorm.DB

func init() {
	//open a db connection
	var err error
	port := os.Getenv("PORT")
	if port != "" {
		Db, err = gorm.Open("postgres", "host=ec2-52-71-85-210.compute-1.amazonaws.com"+
			" port=5432 user=zswecpzsyyjkqz dbname=d1fmiuhjc25u5f"+
			" password=9ad7aeeeacc2ca1ff64f9cb29c7ae37a9127a905f9c3267aeafe90207405886a")
	} else {
		Db, err = gorm.Open("mysql", "root:password@/golang?charset=utf8&parseTime=True&loc=Local")
	}

	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	Db.AutoMigrate(&models.Admin{},
		&models.Product{}, &models.User{}, &models.Vendor{}, &models.Order{}, &models.ProductPurchased{})
}
