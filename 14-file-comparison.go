package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Function to read a file and return its lines
func readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Function to write the comparison results to a file
func writeResultsToFile(outputDir, fileName string, results []string) error {
	// Ensure the directory exists or create it
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	// Create or overwrite the result file
	outputPath := filepath.Join(outputDir, fileName)
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	// Write results to the file
	writer := bufio.NewWriter(file)
	for _, line := range results {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to output file: %v", err)
		}
	}
	writer.Flush()

	return nil
}

func main() {
	// Get file paths from the user
	var filePath1 = "E:\\[0] DOCUMENTS\\CSE687-SPRING-2025\\NewCode\\main.cpp"
	var filePath2 = "E:\\[0] DOCUMENTS\\CSE687-SPRING-2025\\NewCode\\new-main.cpp"
	var outputDir = "E:\\[0] DOCUMENTS\\CSE687-SPRING-2025\\NewCode\\"

	// Read the contents of the files
	file1Lines, err := readFileLines(filePath1)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath1, err)
		return
	}

	file2Lines, err := readFileLines(filePath2)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath2, err)
		return
	}

	// Compare the files line by line
	maxLines := len(file1Lines)
	if len(file2Lines) > maxLines {
		maxLines = len(file2Lines)
	}

	var results []string
	differencesFound := false

	for i := 0; i < maxLines; i++ {
		var file1Line, file2Line string

		if i < len(file1Lines) {
			file1Line = file1Lines[i]
		} else {
			file1Line = "[No Line]"
		}

		if i < len(file2Lines) {
			file2Line = file2Lines[i]
		} else {
			file2Line = "[No Line]"
		}

		if file1Line != file2Line {
			differencesFound = true
			results = append(results, fmt.Sprintf("Difference at line %d:", i+1))
			results = append(results, fmt.Sprintf("  File 1: %s", file1Line))
			results = append(results, fmt.Sprintf("  File 2: %s", file2Line))
		}
	}

	if !differencesFound {
		results = append(results, "The files are identical.")
	} else {
		results = append(results, "Comparison complete. Differences were found.")
	}

	// Write the results to a file
	outputFileName := "ComparisonResults.txt"
	err = writeResultsToFile(outputDir, outputFileName, results)
	if err != nil {
		fmt.Printf("Error writing results to file: %v\n", err)
		return
	}

	fmt.Printf("Comparison results saved to: %s\n", filepath.Join(outputDir, outputFileName))
}
