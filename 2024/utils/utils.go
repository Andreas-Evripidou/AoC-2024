package utils

import (
	"log"
	"os"
	"strings"
)

// ReadInput reads the content of a file and returns it as a string
func ReadInput(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return strings.TrimSpace(string(data))
}
