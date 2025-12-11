package handlers


import (
	"meachine-test/database"
	"meachine-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create product (already done earlier)
func CreateVariantOption(c *gin.Context) {
	var variant models.VariantOption

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
func ListVariantOptios(c *gin.Context) {
	var variants []models.ListVariantOption

	result := database.DB.
		Preload("Variant").
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
