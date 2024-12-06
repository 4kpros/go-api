package utils

import "fmt"

func InterfaceToError(recovery interface{}) (err error) {
	if castedErr, ok := recovery.(error); ok {
		err = castedErr
	} else {
		err = fmt.Errorf("%v", recovery)
	}
	return
}
