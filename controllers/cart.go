package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/models"
)

type AddToCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func AddToCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddToCartRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.MustGet("userID").(uint)

		var cart models.Cart
		if err := config.DB.FirstOrCreate(&cart, models.Cart{UserID: userID}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var item models.CartItem
		if err := config.DB.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&item).Error; err == nil {

			item.Quantity += req.Quantity
			config.DB.Save(&item)
		} else {
			item = models.CartItem{
				CartID:    cart.ID,
				ProductID: req.ProductID,
				Quantity:  req.Quantity,
			}
			config.DB.Create(&item)
		}

		c.JSON(http.StatusOK, gin.H{"message": "added to cart"})
	}
}

func ViewCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(uint)

		var cart models.Cart
		if err := config.DB.Preload("Items").First(&cart, "user_id = ?", userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
			return
		}

		c.JSON(http.StatusOK, cart)
	}
}
