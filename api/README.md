# Optimus Backend

The backend for Optimus is a Go-based API that aggregates telecom industry news.

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Hosting:** GCP Cloud Run
- **News API:** NewsAPI.org

## Setup

### Prerequisites

- Go (>= 1.19)
- Docker
- A NewsAPI key (https://newsapi.org/)
- GCP account for deployment

### Local Development

1. **Get your API Key**  
   Go to https://newsapi.org/

2. **Set up environment variables**  
   Create a `.env` file:

   ```sh
   NEWS_API_KEY=your_api_key_here
   ```

3. **Run the application**

   ```sh
   go run main.go
   ```

4. **Run with Docker**

   ```sh
   docker build -t optimus-backend .
   docker run -p 8080:8080 --env-file .env optimus-backend
   ```

5. **API Endpoints**
   - `GET /` → Health check
   - `GET /news` → Fetch latest telecom news

## API Documentation

### GET /news

Returns curated telecom industry news articles.

#### Response:

```json
{
  "articles": [
    {
      "title": "string",
      "description": "string",
      "url": "string"
    }
  ]
}
```

## Backend Hosting & Deployment

The backend API is hosted on **Google Cloud Platform** using **Google Cloud Run**, a fully managed serverless platform. The API is containerized using **Docker** and pushed to **Google Artifact Registry** before being deployed to Cloud Run. After deployment, Cloud Run provides a publicly accessible URL for the API, which can be used by the frontend or other services.
