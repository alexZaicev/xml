package errors

import "fmt"

type IllegalArgumentError struct {
	baseError
	Arg string
}

func NewIllegalArgumentError(arg, msg string) *IllegalArgumentError {
	return &IllegalArgumentError{
		baseError: newBaseError(
			fmt.Sprintf("illegal argument error: %s %s", arg, msg),
			nil,
		),
		Arg: arg,
	}
}
