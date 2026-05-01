package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FormatJSON reads a JSON file, validates it, and returns a pretty-printed version.
func FormatJSON(filePath string) (string, error) {
	// Step 1: Validate the file path before doing any I/O
	if err := validateJSONFile(filePath); err != nil {
		return "", err
	}

	// Step 2: Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}

	// Step 3: Check for empty file
	if len(strings.TrimSpace(string(data))) == 0 {
		return "", fmt.Errorf("file '%s' is empty — nothing to format", filePath)
	}

	// Step 4: Unmarshal to validate JSON structure
	var parsedJSON interface{}
	if err := json.Unmarshal(data, &parsedJSON); err != nil {
		return "", fmt.Errorf("invalid JSON in '%s': %w", filePath, err)
	}

	// Step 5: Re-marshal with indentation
	prettyData, err := json.MarshalIndent(parsedJSON, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}

	return string(prettyData), nil
}

// validateJSONFile checks that the file exists and has a .json extension.
func validateJSONFile(filePath string) error {
	// Check if file exists
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file not found: '%s'", filePath)
	}
	if err != nil {
		return fmt.Errorf("cannot access file '%s': %w", filePath, err)
	}

	// Don't allow directories
	if info.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", filePath)
	}

	// Warn if not a .json extension
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext != ".json" {
		return fmt.Errorf("expected a .json file, got '%s'", ext)
	}

	return nil
}
