package helpers

import (
	"os"
)

func ReadFileContentToString(path string) (contentStr string, err error) {
	content, errRead := os.ReadFile(path)
	if errRead != nil {
		err = errRead
		return
	}
	contentStr = string(content)
	return
}
