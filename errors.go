package form

import (
	"errors"
)

// Errors is a list of errors that occured when processing a slice of processors.
type Errors []error

// Error implements the error interface.
func (Errors) Error() string {
	return "multiple error occurred"
}

// ErrorMap is a map of all of the keys that experienced errors.
type ErrorMap map[string]error

// Error implements the error interface.
func (ErrorMap) Error() string {
	return "errors occurred during processing form data"
}

// Errors.
var (
	ErrNeedPointer     = errors.New("need pointer to type")
	ErrNeedStruct      = errors.New("need struct type")
	ErrNotInRange      = errors.New("value not in valid range")
	ErrInvalidBoolean  = errors.New("invalid boolean")
	ErrRequiredMissing = errors.New("required field is missing")
	ErrNoMatch         = errors.New("string did not match regex")
)
