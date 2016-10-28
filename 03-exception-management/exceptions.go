package exceptions

import "errors"

//ThrowException returns a 0 value along with an exception (error)
func ThrowException() (int, error) {
	return 0, errors.New("This is an error")
}
