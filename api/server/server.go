package server

import (
	"log"
	"net/http"

	"optimus/newsapi"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Optimus!"})
	})

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
