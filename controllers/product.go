package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/models"
)

type ProductInput struct {
	Name        string
	Description string
	Price       int64
	Stock       int32
	Image       string
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product retrieved successfully", "data": product})
}

func CreateProduct(c *gin.Context) {
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data"})
	}
	product := models.Product{Name: productInput.Name, Price: productInput.Price, Stock: productInput.Stock, Description: productInput.Description, Image: productInput.Image}

	result := config.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func GetAllProducts(c *gin.Context) {

	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the products"})
	}

	c.JSON(http.StatusOK,
		gin.H{"data": products, "message": "Products retrieved successfully"})
}
