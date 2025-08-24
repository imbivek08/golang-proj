package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/models"
	"gorm.io/gorm"
)

type CreateOrderRequest struct {
	Items []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func CreateOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.MustGet("userID").(uint)
		order := models.Order{
			UserID: userID,
		}

		err := config.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&order).Error; err != nil {
				return err
			}

			var total float64
			for _, item := range req.Items {
				var product models.Product
				if err := tx.First(&product, item.ProductID).Error; err != nil {
					return fmt.Errorf("product %d not found", item.ProductID)
				}
				fmt.Println(product.Price)
				subtotal := float64(item.Quantity) * product.Price
				total += subtotal

				orderItem := models.OrderItem{
					OrderID:   order.ID,
					ProductID: product.ID,
					Quantity:  item.Quantity,
					Subtotal:  subtotal,
				}

				if err := tx.Create(&orderItem).Error; err != nil {
					return err
				}
			}

			order.TotalPrice = total
			if err := tx.Save(&order).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, order)
	}
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
