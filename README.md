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

#### type Float32

```go
type Float32 struct {
	Data *float32
}
```

Float32 is a float32 that implements Parser

#### type Float64

```go
type Float64 struct {
	Data *float64
}
```

Float64 is a float64 that implements Parser

#### type FloatBetween

```go
type FloatBetween struct {
	Parser
	Min, Max float64
}
```

FloatBetween is an implementation of Parser that allows a float between the Min
and Max (inclusively)

#### type Int

```go
type Int struct {
	Data *int
}
```

Int is an int that implements Parser

#### type Int16

```go
type Int16 struct {
	Data *int16
}
```

Int16 is an int16 that implements Parser

#### type Int32

```go
type Int32 struct {
	Data *int32
}
```

Int32 is an int32 that implements Parser

#### type Int64

```go
type Int64 struct {
	Data *int64
}
```

Int64 is an int64 that implements Parser

#### type Int8

```go
type Int8 struct {
	Data *int8
}
```

Int8 is an int8 that implements Parser

#### type IntBetween

```go
type IntBetween struct {
	Parser
	Min, Max int64
}
```

IntBetween is an implementation of Parser that allows a signed integer between
the Min and Max (inclusively)

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

#### type RequiredString

```go
type RequiredString String
```

RequiredString is like String except that it returns an error when it gets an
empty string

#### type String

```go
type String struct {
	Data *string
}
```

String is a string that implements Parser

#### type Time

```go
type Time struct {
	Data *time.Time
}
```

Time is a time.Time that implements Parser

#### type TimeFormat

```go
type TimeFormat struct {
	Data   *time.Time
	Format string
}
```

TimeFormat is like a Time except that it allows a specific format to be
specified

#### type Uint

```go
type Uint struct {
	Data *uint
}
```

Uint is a uint that implements Parser

#### type Uint16

```go
type Uint16 struct {
	Data *uint16
}
```

Uint16 is a uint16 that implements Parser

#### type Uint32

```go
type Uint32 struct {
	Data *uint32
}
```

Uint32 is a uint32 that implements Parser

#### type Uint64

```go
type Uint64 struct {
	Data *uint64
}
```

Uint64 is a uint64 that implements Parser

#### type Uint8

```go
type Uint8 struct {
	Data *uint8
}
```

Uint8 is a uint8 that implements Parser

#### type UintBetween

```go
type UintBetween struct {
	Parser
	Min, Max uint64
}
```

UintBetween is an implementation of Parser that allows an unsigned integer
between the Min and Max (inclusively)

#### type UnixTime

```go
type UnixTime struct {
	Data *time.Time
}
```

UnixTime is like time except that it expects the time to be a Unix timestamp
