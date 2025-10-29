package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/smae1993/my_first_go_web_api/modules/user"
)

func initDB() *gorm.DB {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	db.AutoMigrate(&user.User{})
	return db
}

func main() {
	db := initDB()
	userService := user.NewService(db)
	userHandler := user.NewHandler(userService)
	r := gin.Default()
	userHandler.RegisterRoutes(r)
	r.Run(":8080")
}
