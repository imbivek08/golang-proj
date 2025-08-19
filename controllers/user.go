package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/models"
	"github.com/imbivek08/web-service/utils"
)

type UpdateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UpdateUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var input UpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not fould"})
		return
	}
	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Password != "" {
		hashed, _ := utils.HashPassword(input.Password)
		user.Password = string(hashed)
	}
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}

func GetUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id":  user.ID,
		"username": user.Username,
	})

}

func DeleteUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	if err := config.DB.Delete(&models.User{}, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
