package controllers

import (
	"friendship_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListFriendRequests(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Email string `json:"email"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var user models.User
        if err := db.First(&user, "email = ?", req.Email).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        var requests []models.Friend
        db.Preload("Requestor").Where("to_id = ?", user.ID).Find(&requests)

        response := gin.H{"requests": []gin.H{}}
        for _, request := range requests {
            response["requests"] = append(response["requests"].([]gin.H), gin.H{
                "requestor": request.Requestor.Email,
                "status":    request.Status,
            })
        }

        c.JSON(http.StatusOK, response)
    }
}

func ListFriends(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Email string `json:"email"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var user models.User
        if err := db.Preload("Friendships").First(&user, "email = ?", req.Email).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        response := gin.H{"friends": []string{}}
        for _, friend := range user.Friendships {
            response["friends"] = append(response["friends"].([]string), friend.Email)
        }

        c.JSON(http.StatusOK, response)
    }
}

func CommonFriends(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Friends []string `json:"friends"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if len(req.Friends) != 2 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Exactly two friends are required"})
            return
        }

        var users [2]models.User
        for i, email := range req.Friends {
            if err := db.Preload("Friendships").First(&users[i], "email = ?", email).Error; err != nil {
                c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
                return
            }
        }

        commonFriends := []string{}
        for _, friend1 := range users[0].Friendships {
            for _, friend2 := range users[1].Friendships {
                if friend1.ID == friend2.ID {
                    commonFriends = append(commonFriends, friend1.Email)
                }
            }
        }

        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "friends": commonFriends,
        	"count": len(commonFriends),
        })
    }
}

func BlockUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Requestor string `json:"requestor"`
            Block     string `json:"block"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var requestor, blockUser models.User
        if err := db.FirstOrCreate(&requestor, models.User{Email: req.Requestor}).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }
        if err := db.FirstOrCreate(&blockUser, models.User{Email: req.Block}).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }

        if err := db.Model(&requestor).Association("Blocked").Append(&blockUser); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"success": true})
    }
}

