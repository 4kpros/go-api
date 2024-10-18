package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// ReadMultipartFile Reads multipart(form) file and returns the buffer
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

// ReadFileToString Reads the contents of a file into a string.
//
// It returns a pointer to the string for potential performance
// optimization when dealing with large files content.
func ReadFileToString(path string) (*string, error) {
	buffer, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	result := string(buffer)
	return &result, nil
}

func ReadFile(path string) (result []byte, err error) {
	absPath, _ := filepath.Abs(path)
	result, err = os.ReadFile(absPath)
	if err != nil {
		return
	}
	return
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
	initialFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), GenerateRandomAlphaNumeric(20))

	// Create temp file
	absPath, _ := filepath.Abs(path)
	tempFile, err := os.CreateTemp(absPath, initialFileName)
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

	// Retrieve the new file name
	fileName = tempFile.Name()[len(absPath)+1:]
	return
}
