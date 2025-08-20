package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/models"
)

func CreateOrder(c *gin.Context) {
	var orderInput models.Order
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	//totalPrice := orderInput.Quantity*
	order := models.Order{UserID: orderInput.UserID, ProductID: orderInput.ProductID, Quantity: orderInput.Quantity}
	result := config.DB.Create(&order)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully"})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order delete successfully"})
}

func UpdateOrder(c *gin.Context) {

}
