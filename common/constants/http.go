package constants

import "fmt"

var Http500ErrorMessage = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("Error occurred when trying to %s! Please try again later.", message))
}

var Http401InvalidTokenErrorMessage = func() error {
	return fmt.Errorf("%s", "Invalid or expired token! Please enter valid information.")
}

var Http403InvalidPermissionErrorMessage = func() error {
	return fmt.Errorf("%s", "You don't have permission to access this resource! Please enter valid information.")
}

var Http404ErrorMessage = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("%s not found! Please enter valid information.", message))
}

var Http302ErrorMessage = func(message string) error {
	return fmt.Errorf("%s", fmt.Sprintf("This %s already exists! Please enter valid information.", message))
}

const DefaultBodySize = 10000 // 10Kb
