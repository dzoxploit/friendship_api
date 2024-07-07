package main

import (
	"friendship_api/controllers"
	"friendship_api/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        log.Fatal("MYSQL_DSN environment variable not set")
    }

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&models.User{}, &models.Friend{})

   r := gin.Default()

    r.POST("/friend-request", controllers.CreateFriendRequest(db))
    r.POST("/friend-request/respond", controllers.RespondFriendRequest(db))
    r.GET("/friend-requests", controllers.ListFriendRequests(db))
    r.GET("/friends", controllers.ListFriends(db))
    r.GET("/common-friends", controllers.CommonFriends(db))
    r.POST("/block-user", controllers.BlockUser(db))

    r.Run(":8080")
}
