package constants

import "fmt"

var HTTP_500_ERROR_MESSAGE = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("Error occurred when trying to %s! Please try again later.", message))
}

var HTTP_404_ERROR_MESSAGE = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("%s not found! Please enter valid information.", message))
}

var HTTP_302_ERROR_MESSAGE = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("This %s already exists! Please enter valid information.", message))
}
