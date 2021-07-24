package form

import (
	"errors"
	"fmt"
)

// Errors is a list of all form processing errors
type Errors []error

// Error implements the error interface
func (Errors) Error() string {
	return "errors occurred during processing form data"
}

// ErrRequiredMissing is an error returned when a required form value is not
// specified
type ErrRequiredMissing string

// Error implements the error interface
func (ErrRequiredMissing) Error() string {
	return "required value missing"
}

// ErrProcessingFailed is an error describing a failed data processing
type ErrProcessingFailed struct {
	Key string
	Err error
}

// Error implements the error interface
func (e ErrProcessingFailed) Error() string {
	return fmt.Sprintf("error processing key %q: %s", e.Key, e.Err)
}

// Unwrap retrieves the underlying error
func (e ErrProcessingFailed) Unwrap() error {
	return e.Err
}

// Errors
var (
	ErrNeedPointer = errors.New("need pointer to type")
	ErrNeedStruct  = errors.New("need struct type")
	ErrNotInRange  = errors.New("value not in valid range")
)
