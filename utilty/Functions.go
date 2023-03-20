package utilty

import "errors"

func ErrorReturner[T interface{}](message string) (T, error) {
	return *new(T), errors.New(message)
}
