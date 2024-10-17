package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func ReadMultipartFile(file multipart.File) ([]byte, error) {
	defer func(File multipart.File) {
		err := File.Close()
		if err != nil {
			return
		}
	}(file)
	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func ReadFile(path string) (result []byte, err error) {
	absPath, _ := filepath.Abs(path)
	result, err = os.ReadFile(absPath)
	if err != nil {
		return
	}
	return
}

// Reads the entire contents of a file into a string.
//
// It returns a pointer to the string for potential performance
// optimization when dealing with large files content.
func ReadFileToString(path string) (*string, error) {
	absPath, _ := filepath.Abs(path)
	content, errRead := os.ReadFile(absPath)
	if errRead != nil {
		return nil, errRead
	}
	result := string(content)
	return &result, nil
}

func DeleteFile(path string) (isDeleted bool, err error) {
	absPath, _ := filepath.Abs(path)
	err = os.Remove(absPath)
	if err != nil {
		isDeleted = false
		return
	}
	isDeleted = true
	return
}

func SaveFile(buffer []byte, path string) (fileName string, err error) {
	// Encode file name
	fileName = fmt.Sprintf("%d_%s", time.Now().Unix(), GenerateRandomAlphaNumeric(20))

	// Create temp file
	absPath, _ := filepath.Abs(path)
	tempFile, err := os.CreateTemp(absPath, fileName)
	if err != nil {
		return
	}
	defer func(tempFile *os.File) {
		err = tempFile.Close()
		if err != nil {
			return
		}
	}(tempFile)

	// Write butter to the file
	_, err = tempFile.Write(buffer)
	if err != nil {
		return
	}
	return
}
