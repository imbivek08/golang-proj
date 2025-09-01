package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/models"
	"gorm.io/gorm"
)

func CheckoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(uint)

		err := config.DB.Transaction(func(tx *gorm.DB) error {
			// 1. Get all cart items for user
			var cartItems []models.CartItem
			if err := tx.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
				return err
			}
			if len(cartItems) == 0 {
				return fmt.Errorf("cart is empty")
			}

			// 2. Create a new order
			order := models.Order{
				UserID: userID,
			}
			if err := tx.Create(&order).Error; err != nil {
				return err
			}

			// 3. Copy cart items -> order_items
			var total float64
			for _, cartItem := range cartItems {
				var product models.Product
				if err := tx.First(&product, cartItem.ProductID).Error; err != nil {
					return fmt.Errorf("product %d not found", cartItem.ProductID)
				}

				subtotal := float64(cartItem.Quantity) * product.Price
				total += subtotal

				orderItem := models.OrderItem{
					OrderID:   order.ID,
					ProductID: cartItem.ProductID,
					Quantity:  int(cartItem.Quantity),
					Subtotal:  subtotal,
				}
				if err := tx.Create(&orderItem).Error; err != nil {
					return err
				}
			}

			// 4. Update order with total price
			order.TotalPrice = total
			if err := tx.Save(&order).Error; err != nil {
				return err
			}

			// 5. Clear user cart
			if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Checkout successful, order created"})
	}
}
