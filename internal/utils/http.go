package utils

import (
	"fmt"
	"io"
	"net/http"
)

// FetchData makes an HTTP GET request to the provided URL and returns the body as a string.
func FetchData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		// Wrap errors to provide context
		return "", fmt.Errorf("network error occurred: %w", err)
	}
	// Important: Always defer closing the response body to prevent memory leaks
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("received non-ok HTTP status code: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
