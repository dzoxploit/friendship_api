package controllers

import (
	"friendship_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFriendRequest(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Requestor string `json:"requestor"`
            To        string `json:"to"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var requestor, to models.User
        if err := db.FirstOrCreate(&requestor, models.User{Email: req.Requestor}).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }
        if err := db.FirstOrCreate(&to, models.User{Email: req.To}).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }

        // Check if user is blocked
		var blockedUsers []models.User
		db.Model(&requestor).Association("Blocked").Find(&blockedUsers, "email = ?", req.To)
		if len(blockedUsers) > 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "User is blocked"})
			return
		}

        friendRequest := models.Friend{
            RequestorID: requestor.ID,
            ToID:        to.ID,
            Status:      "pending",
        }

        if err := db.Create(&friendRequest).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"success": true})
    }
}

func RespondFriendRequest(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req struct {
            Requestor string `json:"requestor"`
            To        string `json:"to"`
            Action    string `json:"action"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var requestor, to models.User
        if err := db.First(&requestor, "email = ?", req.Requestor).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Requestor not found"})
            return
        }
        if err := db.First(&to, "email = ?", req.To).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Recipient not found"})
            return
        }

        var friendRequest models.Friend
        if err := db.Where("requestor_id = ? AND to_id = ? AND status = 'pending'", requestor.ID, to.ID).First(&friendRequest).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Friend request not found"})
            return
        }

        if req.Action == "accept" {
            friendRequest.Status = "accepted"
            db.Model(&requestor).Association("Friendships").Append(&to)
            db.Model(&to).Association("Friendships").Append(&requestor)
        } else {
            friendRequest.Status = "rejected"
        }

        if err := db.Save(&friendRequest).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"success": true})
    }
}
