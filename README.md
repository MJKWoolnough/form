# form
--
    import "vimagination.zapto.org/form"

Package form provides an easy to use way to parse form values from an HTTP
request into a struct

## Usage

```go
var (
	ErrNeedPointer    = errors.New("need pointer to type")
	ErrNeedStruct     = errors.New("need struct type")
	ErrNotInRange     = errors.New("value not in valid range")
	ErrInvalidBoolean = errors.New("invalid boolean")
)
```
Errors

#### func  Process

```go
func Process(r *http.Request, fv interface{}) error
```
Process parses the form data from the request into the passed value, which must
be a pointer to a struct.

Form keys are assumed to be the field names unless a 'form' tag is provided with
an alternate name, for example, in the following struct, the int is parse with
key 'A' and the bool is parsed with key 'C'.

type Example struct {
    A int
    B bool `form:"C"`
}

Two options can be added to the form tag to modify the processing. The 'post'
option forces the processer to parse a value from the PostForm field of the
Request, and the 'required' option will have an error thrown if the key in not
set.

Number types can also have minimums and maximums checked during processing by
setting the min and max tags accordingly.

Lastly, a custom data processor can be specified by attaching a method to the
field type with the following specification:

ParseForm([]string) error

#### type ErrProcessingFailed

```go
type ErrProcessingFailed struct {
	Key string
	Err error
}
```

ErrProcessingFailed is an error describing a failed data processing

#### func (ErrProcessingFailed) Error

```go
func (e ErrProcessingFailed) Error() string
```
Error implements the error interface

#### func (ErrProcessingFailed) Unwrap

```go
func (e ErrProcessingFailed) Unwrap() error
```
Unwrap retrieves the underlying error

#### type ErrRequiredMissing

```go
type ErrRequiredMissing string
```

ErrRequiredMissing is an error returned when a required form value is not
specified

#### func (ErrRequiredMissing) Error

```go
func (ErrRequiredMissing) Error() string
```
Error implements the error interface

#### type Errors

```go
type Errors []error
```

Errors is a list of all form processing errors

#### func (Errors) Error

```go
func (Errors) Error() string
```
Error implements the error interface
