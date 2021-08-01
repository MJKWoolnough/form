# form
--
    import "vimagination.zapto.org/form"

Package form provides an easy to use way to parse form values from an HTTP
request into a struct

## Usage

```go
var (
	ErrNeedPointer     = errors.New("need pointer to type")
	ErrNeedStruct      = errors.New("need struct type")
	ErrNotInRange      = errors.New("value not in valid range")
	ErrInvalidBoolean  = errors.New("invalid boolean")
	ErrRequiredMissing = errors.New("required field is missing")
	ErrNoMatch         = errors.New("string did not match regex")
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
setting the 'min' and 'max' tags accordingly.

In a similar vein, string types can utilise the 'regex' tag to set a regular
expression to be matched against.

Anonymous structs are traversed, but will not override more local fields.

Slices of basic types can be processed, and errors returned from any such
processing will be of the Errors type, which each indexed entry corresponding to
the index of the processed data.

Pointers to basic types can also be processed, with the type being allocated
even if an error occurs.

Lastly, a custom data processor can be specified by attaching a method to the
field type with the following specification:

ParseForm([]string) error

#### type ErrorMap

```go
type ErrorMap map[string]error
```

ErrorMap is a map of all of the keys that experienced errors

#### func (ErrorMap) Error

```go
func (ErrorMap) Error() string
```
Error implements the error interface

#### type Errors

```go
type Errors []error
```

Errors is a list of errors that occured when processing a slice of processors

#### func (Errors) Error

```go
func (Errors) Error() string
```
Error implements the error interface
