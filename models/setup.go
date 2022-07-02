package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL

)

var DB *gorm.DB

func DBconnectDatabase() {
	//  db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "user:testpass/billing_client_prepration")
	//database, err := gorm.Open("sqlite3", "test.db")
	dsn := "testing:testpass@tcp(127.0.0.1:3306)/billing_client_prepration?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//defer db.Close()
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&EmployeeWork{})

	DB = db
}
