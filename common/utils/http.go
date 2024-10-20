package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// HttpGet Fetches url and return response
func HttpGet(url string, response any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}
