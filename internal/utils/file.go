package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// OrganizeFolder iterates over files in the target directory and moves them into categorized subfolders.
func OrganizeFolder(targetDir string) error {
	// Read all entries in the directory
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return fmt.Errorf("failed to read directory '%s': %w", targetDir, err)
	}

	for _, entry := range entries {
		// Ignore nested directories, we only want to organize root-level files
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		ext := strings.ToLower(filepath.Ext(fileName))
		
		// Skip files without extensions (like dotfiles or binaries)
		if ext == "" {
			continue 
		}

		// Determine the subfolder based on the extension
		categoryName := getCategoryForExtension(ext)
		destFolderPath := filepath.Join(targetDir, categoryName)

		// Create the subfolder if it does not exist yet (0755 permissions is standard)
		if err := os.MkdirAll(destFolderPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory '%s': %w", destFolderPath, err)
		}

		// Move the file into the new directory
		oldPath := filepath.Join(targetDir, fileName)
		newPath := filepath.Join(destFolderPath, fileName)
		
		// os.Rename acts as a move command
		if err := os.Rename(oldPath, newPath); err != nil {
			return fmt.Errorf("failed to move file '%s' to '%s': %w", fileName, newPath, err)
		}
	}

	return nil
}

// getCategoryForExtension acts as a router mapping file extensions to human-readable folder names.
func getCategoryForExtension(ext string) string {
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".svg", ".webp", ".ico":
		return "Images"
	case ".pdf", ".doc", ".docx", ".txt", ".md", ".csv":
		return "Documents"
	case ".mp4", ".mkv", ".avi", ".mov":
		return "Videos"
	case ".mp3", ".wav", ".flac":
		return "Audio"
	case ".zip", ".tar", ".gz", ".rar":
		return "Archives"
	case ".go", ".js", ".ts", ".py", ".java", ".json", ".html", ".css":
		return "Code"
	default:
		return "Others"
	}
}
