package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const mockLocationResponse = `{
	"count": 2,
	"next": "https://pokeapi.co",
	"previous": null,
	"results": [
		{"name": "kanto-route-1-area", "url": "https://pokeapi.co"},
		{"name": "kanto-route-2-area", "url": "https://pokeapi.co"}
	]
}`

func TestNewClient(t *testing.T) {
	timeout := 2 * time.Second
	cacheInterval := 5 * time.Minute

	client := NewClient(timeout, cacheInterval)

	if client.cache == nil {
		t.Error("expected cache to be initialized, got nil")
	}
	if client.httpClient.Timeout != timeout {
		t.Errorf("expected HTTP client timeout to be %v, got %v", timeout, client.httpClient.Timeout)
	}
}

func TestGetLocationAreas_SuccessAndCache(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockLocationResponse))
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)

	targetURL := server.URL

	res1, err := client.GetLocationAreas(&targetURL)
	if err != nil {
		t.Fatalf("unexpected error on first call: %v", err)
	}

	if res1.Count != 2 {
		t.Errorf("expected count 2, got %d", res1.Count)
	}
	if len(res1.Results) != 2 || res1.Results[0].Name != "kanto-route-1-area" {
		t.Errorf("unexpected results content: %v", res1.Results)
	}
	if callCount != 1 {
		t.Errorf("expected exactly 1 network call, got %d", callCount)
	}

	res2, err := client.GetLocationAreas(&targetURL)
	if err != nil {
		t.Fatalf("unexpected error on second call: %v", err)
	}

	if res2.Count != res1.Count {
		t.Errorf("cached result count mismatch: got %d, expected %d", res2.Count, res1.Count)
	}

	// If caching is operational, the network call count should STILL be 1
	if callCount != 1 {
		t.Errorf("cache missed! Expected network calls to stay at 1, but hit the server %d times", callCount)
	}
}

func TestGetLocationAreas_HttpError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Return 404 Not Found
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	targetURL := server.URL

	_, err := client.GetLocationAreas(&targetURL)
	if err == nil {
		t.Fatal("expected an error due to 404 status code, got nil")
	}

	expectedErrSnippet := "request failed: 404 Not Found"
	if err.Error() != expectedErrSnippet {
		t.Errorf("expected error message containing %q, got %q", expectedErrSnippet, err.Error())
	}
}

func TestGetLocationAreas_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{ "count": "not-an-integer-value" }`)) // Bad JSON data types
	}))
	defer server.Close()

	client := NewClient(5*time.Second, 5*time.Minute)
	targetURL := server.URL

	_, err := client.GetLocationAreas(&targetURL)
	if err == nil {
		t.Fatal("expected a JSON decoding error, got nil")
	}
}

func TestGetLocationAreas_NetworkFailure(t *testing.T) {
	client := NewClient(5*time.Second, 5*time.Minute)

	// Provide a garbage URL that will immediately fail lookups
	brokenURL := "http://localhost:99999/invalid-url-path"

	_, err := client.GetLocationAreas(&brokenURL)
	if err == nil {
		t.Fatal("expected a network routing error, got nil")
	}
}
