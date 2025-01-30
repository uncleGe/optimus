package newsapi

import (
	"strings"
	"testing"
)

func TestQueryString(t *testing.T) {
	apiKey := "test-api-key"
	url := BuildNewsQuery(apiKey)

	if !strings.Contains(url, BaseQuery) {
		t.Errorf("Expected query to contain BaseQuery, but got: %s", url)
	}
	if !strings.Contains(url, apiKey) {
		t.Errorf("Expected query to contain API key, but got: %s", url)
	}
}
