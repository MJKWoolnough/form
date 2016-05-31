# form
--
    import "github.com/MJKWoolnough/form"

Package form provides an easy to use way to parse form values from an HTTP
request into a struct or other data object

## Usage

#### func  Parse

```go
func Parse(p ParserLister, data url.Values) error
```
Parse parses the given url.Values into the type given

#### func  ParseValue

```go
func ParseValue(name string, value Parser, data url.Values) error
```
ParseValue parses a single values

#### type Bool

```go
type Bool struct {
	Data *bool
}
```

Bool is a bool that implements Parser

#### func (Bool) Parse

```go
func (b Bool) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Empty

```go
type Empty struct{}
```

Empty is returned when a field contains only an empty string and that is not
allowed

#### func (Empty) Error

```go
func (Empty) Error() string
```

#### type Errors

```go
type Errors map[string]error
```

Errors is an error type that is a map of other errors

#### func (Errors) Error

```go
func (Errors) Error() string
```

#### type Float32

```go
type Float32 struct {
	Data *float32
}
```

Float32 is a float32 that implements Parser

#### func (Float32) Parse

```go
func (f Float32) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Float64

```go
type Float64 struct {
	Data *float64
}
```

Float64 is a float64 that implements Parser

#### func (Float64) Parse

```go
func (f Float64) Parse(d []string) error
```
Parse is an implementation of Parser

#### type FloatBetween

```go
type FloatBetween struct {
	Parser
	Min, Max float64
}
```

FloatBetween is an implementation of Parser that allows a float between the Min
and Max (inclusively)

#### func (FloatBetween) Parse

```go
func (f FloatBetween) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Int

```go
type Int struct {
	Data *int
}
```

Int is an int that implements Parser

#### func (Int) Parse

```go
func (i Int) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Int16

```go
type Int16 struct {
	Data *int16
}
```

Int16 is an int16 that implements Parser

#### func (Int16) Parse

```go
func (i Int16) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Int32

```go
type Int32 struct {
	Data *int32
}
```

Int32 is an int32 that implements Parser

#### func (Int32) Parse

```go
func (i Int32) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Int64

```go
type Int64 struct {
	Data *int64
}
```

Int64 is an int64 that implements Parser

#### func (Int64) Parse

```go
func (i Int64) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Int8

```go
type Int8 struct {
	Data *int8
}
```

Int8 is an int8 that implements Parser

#### func (Int8) Parse

```go
func (i Int8) Parse(d []string) error
```
Parse is an implementation of Parser

#### type IntBetween

```go
type IntBetween struct {
	Parser
	Min, Max int64
}
```

IntBetween is an implementation of Parser that allows a signed integer between
the Min and Max (inclusively)

#### func (IntBetween) Parse

```go
func (i IntBetween) Parse(d []string) error
```
Parse is an implementation of Parser

#### type NoRegexMatch

```go
type NoRegexMatch string
```

NoRegexMatch is an error returned from RegexString.Parse when the given string
does not match the given regular expression

#### func (NoRegexMatch) Error

```go
func (NoRegexMatch) Error() string
```

#### type OutsideBounds

```go
type OutsideBounds string
```

OutsideBounds is an error returned from a Parser when the parsed value falls
outside of the range given as an acceptable value

#### func (OutsideBounds) Error

```go
func (OutsideBounds) Error() string
```

#### type Parser

```go
type Parser interface {
	Parse([]string) error
}
```

Parser is an interface used to to parse a specfic type

#### type ParserList

```go
type ParserList map[string]Parser
```

ParserList is a simple implementation of a parserLister that simply returns
itself

#### func (ParserList) ParserList

```go
func (p ParserList) ParserList() ParserList
```
ParserList is an implementation of parserLister

#### type ParserLister

```go
type ParserLister interface {
	ParserList() ParserList
}
```

ParserLister is the main interface for this package. The single method
ParserList returns a ParserList map of field names to Parser's

#### type RegexString

```go
type RegexString struct {
	Data  *string
	Regex *regexp.Regexp
}
```

RegexString is an implementation of Parser that matches a string again the given
regular expression

#### func (RegexString) Parse

```go
func (r RegexString) Parse(d []string) error
```
Parse is an implementation of Parser

#### type RequiredString

```go
type RequiredString String
```

RequiredString is like String except that it returns an error when it gets an
empty string

#### func (RequiredString) Parse

```go
func (r RequiredString) Parse(d []string) error
```
Parse is an implementation of Parser

#### type String

```go
type String struct {
	Data *string
}
```

String is a string that implements Parser

#### func (String) Parse

```go
func (s String) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Time

```go
type Time struct {
	Data *time.Time
}
```

Time is a time.Time that implements Parser

#### func (Time) Parse

```go
func (t Time) Parse(d []string) error
```
Parse is an implementation of Parser

#### type TimeFormat

```go
type TimeFormat struct {
	Data   *time.Time
	Format string
}
```

TimeFormat is like a Time except that it allows a specific format to be
specified

#### func (TimeFormat) Parse

```go
func (t TimeFormat) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Uint

```go
type Uint struct {
	Data *uint
}
```

Uint is a uint that implements Parser

#### func (Uint) Parse

```go
func (u Uint) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Uint16

```go
type Uint16 struct {
	Data *uint16
}
```

Uint16 is a uint16 that implements Parser

#### func (Uint16) Parse

```go
func (u Uint16) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Uint32

```go
type Uint32 struct {
	Data *uint32
}
```

Uint32 is a uint32 that implements Parser

#### func (Uint32) Parse

```go
func (u Uint32) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Uint64

```go
type Uint64 struct {
	Data *uint64
}
```

Uint64 is a uint64 that implements Parser

#### func (Uint64) Parse

```go
func (u Uint64) Parse(d []string) error
```
Parse is an implementation of Parser

#### type Uint8

```go
type Uint8 struct {
	Data *uint8
}
```

Uint8 is a uint8 that implements Parser

#### func (Uint8) Parse

```go
func (u Uint8) Parse(d []string) error
```
Parse is an implementation of Parser

#### type UintBetween

```go
type UintBetween struct {
	Parser
	Min, Max uint64
}
```

UintBetween is an implementation of Parser that allows an unsigned integer
between the Min and Max (inclusively)

#### func (UintBetween) Parse

```go
func (u UintBetween) Parse(d []string) error
```
Parse is an implementation of Parser

#### type UnixTime

```go
type UnixTime struct {
	Data *time.Time
}
```

UnixTime is like time except that it expects the time to be a Unix timestamp

#### func (UnixTime) Parse

```go
func (u UnixTime) Parse(d []string) error
```
Parse is an implementation of Parser

#### type UnknownFormat

```go
type UnknownFormat string
```

UnknownFormat is an error returned from Time.Parse when it cannot determine the
time format of the given string

#### func (UnknownFormat) Error

```go
func (UnknownFormat) Error() string
```
