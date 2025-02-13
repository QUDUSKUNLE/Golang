package base

import (
	"encoding/base64"
	"os"
)

// ConvertFileToBase64 reads the file at the given path and returns its contents as a base64 encoded string
func ConvertFileToBase64(filePath string) (string, error) {
	// Read the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Encode the file data to base64
	base64Data := base64.StdEncoding.EncodeToString(fileData)
	return base64Data, nil
}
