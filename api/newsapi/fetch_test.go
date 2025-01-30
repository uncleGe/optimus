package newsapi

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

var mockHTTPGet = func(url string) (*http.Response, error) {
	mockJSON := `{"articles":[{"title":"Test Title","description":"Test Description","url":"http://example.com"}]}`
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(mockJSON))),
	}, nil
}

func TestFetchNews(t *testing.T) {
	originalHTTPGet := httpGet
	defer func() { httpGet = originalHTTPGet }()
	httpGet = mockHTTPGet

	dummyAPIKey := "test-api-key"
	articles := FetchNews(dummyAPIKey)

	if len(articles) != 1 {
		t.Fatalf("Expected 1 article, got %d", len(articles))
	}
	if articles[0].Title != "Test Title" {
		t.Errorf("Expected title 'Test Title', got '%s'", articles[0].Title)
	}
	if articles[0].Description != "Test Description" {
		t.Errorf("Expected description 'Test Description', got '%s'", articles[0].Description)
	}
	if articles[0].URL != "http://example.com" {
		t.Errorf("Expected URL 'http://example.com', got '%s'", articles[0].URL)
	}
}
