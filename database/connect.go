package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	var err error
	const MYSQL = "admin:d2laksana@tcp(127.0.0.1:3306)/iot?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

}
