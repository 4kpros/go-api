package utils

import (
	"fmt"
	"regexp"
)

// ExtractSQLState extracts the first occurrence of text enclosed in parentheses from the input string.
// It is useful for parsing error messages or logs to retrieve specific details like SQL state codes.
//
// Example:
//
//	input := `ERROR: duplicate key value violates unique constraint "uni_roles_name" (SQLSTATE 23505)`
//	result, err := extractSQLState(input)
//	if err != nil {
//	    log.Println("Error:", err)
//	} else {
//	    fmt.Println(result) // Output: 23505
//
// Parameters:
//   - input (string): The string to parse for text inside parentheses.
//
// Returns:
//   - (string): The extracted text from the parentheses if found.
//   - (error): An error if no matching text is found.
func ExtractSQLState(input string) (string, error) {
	reg := regexp.MustCompile(`\((SQLSTATE \d+)\)`)
	matches := reg.FindStringSubmatch(input)
	if len(matches) > 0 {
		return matches[0], nil
	}

	return "", fmt.Errorf("no matching text is found")
}

func ListIntToString(list []int64) string {
	if list == nil {
		return ""
	}
	listLen := len(list)
	if listLen < 1 {
		return ""
	}

	result := ""
	for index, id := range list {
		if index < listLen-1 {
			result += fmt.Sprintf(" %d,", id)
		} else {
			result += fmt.Sprintf(" %d ", id)
		}
	}
	return result
}
