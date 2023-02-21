package utilities

import "errors"

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrSomethingWentWrong = errors.New("something went wrong")
)
