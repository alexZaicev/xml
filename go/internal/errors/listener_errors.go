package errors

import (
	"fmt"
)

type Position interface {
	GetLine() int
	GetColumn() int
}

type ListenerError struct {
	baseError
	Line, Column int
}

func NewListenerError(msg string) *ListenerError {
	return &ListenerError{
		baseError: newBaseError(
			fmt.Sprintf("listener error: %s", msg),
			nil,
		),
	}
}

func NewListenerErrorFromPosition(msg string, position Position) *ListenerError {
	return &ListenerError{
		baseError: newBaseError(
			fmt.Sprintf("listener error: line %d:%d %s", position.GetLine(), position.GetColumn(), msg),
			nil,
		),
		Line:   position.GetLine(),
		Column: position.GetColumn(),
	}
}
