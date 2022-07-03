package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnectDatabase() {
	dsn := "root:P@$$w0rd@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Work{})

	DB = db
}

func (Client) TableName() string {
	return "client"
}

func (Employee) TableName() string {
	return "employee"
}

func (Work) TableName() string {
	return "work"
}
