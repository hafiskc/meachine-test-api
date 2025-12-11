package main

import (
	"meachine-test/database"
	"meachine-test/handlers"
	"time"

	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

func main() {
	// Initialize DB
	database.InitDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Angular dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.ListProducts) //
	r.POST("/variants", handlers.CreateVariant)
	r.GET("/variants", handlers.ListVariants)
	r.POST("/variantoptions", handlers.CreateVariantOption)
	r.GET("/variantoptions", handlers.ListVariantOptios)
	r.Run(":8086") // custom port
}
