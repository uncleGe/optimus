package newsapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type NewsResponse struct {
	Articles []Article `json:"articles"`
}

// Package-level variable to allow HTTP request mocking in tests
var httpGet = http.Get

func LoadAPIKey() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Using system environment variables.")
	}
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("NEWS_API_KEY is not set")
	}
	return apiKey, nil
}

func BuildNewsQuery(apiKey string) string {
	fromDate := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	return fmt.Sprintf(
		"https://newsapi.org/v2/everything?q=%s&from=%s&apiKey=%s",
		BaseQuery, // from query.go
		fromDate,
		apiKey,
	)
}

func FetchNews(apiKey string) []Article {
	if apiKey == "" {
		log.Println("Error: API key is required")
		return nil
	}

	url := BuildNewsQuery(apiKey)
	resp, err := httpGet(url)
	if err != nil {
		log.Println("Error fetching news:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("News API returned status: %d\n", resp.StatusCode)
		return nil
	}

	var result NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding news response:", err)
		return nil
	}

	return result.Articles
}
