package database

import (
	"fmt"
	"log"

	"github.com/vanthang24803/api-ecommerce/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     string = util.GetEnv("DB_HOST")
	port     string = util.GetEnv("DB_PORT")
	username string = util.GetEnv("DB_USERNAME")
	password string = util.GetEnv("DB_PASSWORD")
	name     string = util.GetEnv("DB_NAME")
)

var db *gorm.DB

func init() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, name)

	var err error
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully! ✔️")
}

func GetDb() *gorm.DB {
	return db
}
