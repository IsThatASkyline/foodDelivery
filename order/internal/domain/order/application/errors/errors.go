package errors

import "errors"

var ErrEntityNotFound = errors.New("entity not found")

var ErrItemsMustBeMoreThanZero = errors.New("items must be more than zero")
