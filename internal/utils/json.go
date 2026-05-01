package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// FormatJSON reads an unformatted JSON file and returns a pretty-printed version.
func FormatJSON(filePath string) (string, error) {
	// 1. Read the file into memory
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}

	// 2. Unmarshal into a generic interface to ensure it's valid JSON
	var parsedJSON interface{}
	if err := json.Unmarshal(data, &parsedJSON); err != nil {
		return "", fmt.Errorf("invalid JSON in file: %w", err)
	}

	// 3. Re-marshal it with indentation (2 spaces is standard)
	prettyData, err := json.MarshalIndent(parsedJSON, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to pretty-print JSON: %w", err)
	}

	return string(prettyData), nil
}
