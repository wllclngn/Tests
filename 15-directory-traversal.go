package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Specify the root directory you want to scan
	rootDir := "E:\\[2] MIXES & PODCASTS" // Change this to the directory you want to scan
	outputFile := "E:\\[0] DOCUMENTS\\personal\\personal\\PROGRAMMING\\[02] GO\\directory_structure.md"

	// Create or open the Markdown file
	var file *os.File
	var err error
	existingPaths := make(map[string]bool)

	if _, err = os.Stat(outputFile); err == nil {
		// If the file exists, read its contents to avoid duplications
		file, err = os.OpenFile(outputFile, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("Failed to open existing file: %v\n", err)
			return
		}
		defer file.Close()

		// Load existing paths into a map
		existingPaths = loadExistingPaths(outputFile)
	} else {
		// If the file doesn't exist, create it
		file, err = os.Create(outputFile)
		if err != nil {
			fmt.Printf("Failed to create output file: %v\n", err)
			return
		}
		defer file.Close()

		// Write the Markdown header
		file.WriteString("# Directory Listing\n\n")
		file.WriteString(fmt.Sprintf("Root Directory: `%s`\n\n", rootDir))
	}

	// Walk through the directory
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Indent subfolders based on depth
		relativePath, _ := filepath.Rel(rootDir, path)
		depth := strings.Count(relativePath, string(os.PathSeparator))
		indent := strings.Repeat("  ", depth)

		// Avoid duplicating entries
		if _, exists := existingPaths[relativePath]; exists {
			return nil
		}

		// Write folder or file details
		if info.IsDir() {
			file.WriteString(fmt.Sprintf("%s- **%s**\n", indent, info.Name()))
		} else {
			ext := filepath.Ext(info.Name())
			file.WriteString(fmt.Sprintf("%s  - %s (.%s)\n", indent, info.Name(), strings.TrimPrefix(ext, ".")))
		}

		// Mark path as written
		existingPaths[relativePath] = true
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		return
	}

	fmt.Printf("Directory structure updated in %s\n", outputFile)
}

// loadExistingPaths reads the existing Markdown file and extracts paths to avoid duplicates
func loadExistingPaths(filePath string) map[string]bool {
	paths := make(map[string]bool)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to read existing file: %v\n", err)
		return paths
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Extract relative paths from Markdown (ignoring headers and formatting)
		if strings.HasPrefix(line, "- **") || strings.HasPrefix(line, "  - ") {
			trimmedLine := strings.TrimSpace(line)
			// Extract the name of the file/folder
			entry := strings.Split(trimmedLine, " ")[1]
			entry = strings.Trim(entry, "**")
			paths[entry] = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return paths
}
