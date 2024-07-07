package main

import (
	"friendship_api/controllers"
	"friendship_api/models"
	"log"
	"os"

	"github.com/gin-contrib/cors"
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

   // Apply CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    r.POST("/friend-request", controllers.CreateFriendRequest(db))
    r.POST("/friend-request/respond", controllers.RespondFriendRequest(db))
    r.GET("/friend-requests", controllers.ListFriendRequests(db))
    r.GET("/friends", controllers.ListFriends(db))
    r.GET("/common-friends", controllers.CommonFriends(db))
    r.POST("/block-user", controllers.BlockUser(db))

    r.Run(":8080")
}
