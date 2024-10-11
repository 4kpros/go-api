package helpers

import (
	"os"
)

// Reads the entire file at the specified path into a string.
func ReadFileContentToString(path string) (string, error) {
	content, errRead := os.ReadFile(path)
	if errRead != nil {
		return "", errRead
	}
	return string(content), nil
}
