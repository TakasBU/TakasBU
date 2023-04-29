package databases

import (
	"fmt"
	"log"

	"github.com/TakasBU/TakasBU/initializers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	var err2 error
	dsn := config.DBUserName + ":" + config.DBUserPassword + "@tcp" + "(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err2 != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	return db
}
