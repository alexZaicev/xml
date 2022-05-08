package errors

// baseError is a generic error providing 'error' functionality to OData typed errors.
type baseError struct {
	Err error
	Msg string
}

// NewBaseError create a new instance of baseError.
func newBaseError(msg string, err error) baseError {
	return baseError{
		Err: err,
		Msg: msg,
	}
}

// Error function returns error message.
func (err *baseError) Error() string {
	return err.Msg
}

// Unwrap allows baseError and any structs that embed it to be used with the
// error wrapping utilities introduced in go 1.13.
func (err *baseError) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.Err
}
