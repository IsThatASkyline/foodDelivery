package errors

import "errors"

var ErrEntityNotFound = errors.New("entity not found")

var ErrCannotMakePayment = errors.New("cannot make payment")

var ErrCannotChangeOrderInThisStatus = errors.New("cannot change order in this status")
