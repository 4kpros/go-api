package helpers

import (
	"os"
)

// Reads the entire contents of a file into a string.
//
// It returns a pointer to the string for potential performance
// optimization when dealing with large files content.
func ReadFileContentToString(path string) (*string, error) {
	content, errRead := os.ReadFile(path)
	if errRead != nil {
		return nil, errRead
	}
	var result = string(content)
	return &result, nil
}
