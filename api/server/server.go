package server

import (
	"log"
	"net/http"

	"optimus/newsapi"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// Enable CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // Allow all origins
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Optimus!"})
	})

	// News route
	r.GET("/news", func(c *gin.Context) {
		apiKey, err := newsapi.LoadAPIKey()
		if err != nil {
			log.Println("Error loading API key:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing API key"})
			return
		}

		articles := newsapi.FetchNews(apiKey)
		if articles == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
			return
		}
		c.JSON(http.StatusOK, articles)
	})

	r.Run(":8080")
}
