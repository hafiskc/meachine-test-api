package handlers

import (
	"meachine-test/database"
	"meachine-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create product (already done earlier)
func CreateVariant(c *gin.Context) {
	var variant models.Variant

	if err := c.ShouldBindJSON(&variant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&variant)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    variant,
	})
}

// 👇 NEW: List all products
func ListVariants(c *gin.Context) {
	var variants []models.ListVariant

	result := database.DB.
		Preload("Product").
		Find(&variants)


	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(variants),
		"data":  variants,
	})
}
