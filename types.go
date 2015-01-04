package form

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Bool is a bool that implements Parser
type Bool struct {
	Data *bool
}

// Parse is an implementation of Parser
func (b Bool) Parse(d []string) error {
	switch strings.ToLower(d[0]) {
	case "on", "yes", "y":
		*b.Data = true
	case "off", "no", "n":
		*b.Data = false
	default:
		c, err := strconv.ParseBool(d[0])
		*b.Data = c
		return err
	}
	return nil
}

// Int is an int that implements Parser
type Int struct {
	Data *int
}

// Parse is an implementation of Parser
func (i Int) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 0)
	*i.Data = int(n)
	return err
}

// Int8 is an int8 that implements Parser
type Int8 struct {
	Data *int8
}

// Parse is an implementation of Parser
func (i Int8) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 8)
	*i.Data = int8(n)
	return err
}

// Int16 is an int16 that implements Parser
type Int16 struct {
	Data *int16
}

// Parse is an implementation of Parser
func (i Int16) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 16)
	*i.Data = int16(n)
	return err
}

// Int32 is an int32 that implements Parser
type Int32 struct {
	Data *int32
}

// Parse is an implementation of Parser
func (i Int32) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 32)
	*i.Data = int32(n)
	return err
}

// Int64 is an int64 that implements Parser
type Int64 struct {
	Data *int64
}

// Parse is an implementation of Parser
func (i Int64) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 64)
	*i.Data = n
	return err
}

// IntBetween is an implementation of Parser that allows a signed integer
// between the Min and Max (inclusively)
type IntBetween struct {
	Parser
	Min, Max int64
}

// Parse is an implementation of Parser
func (i IntBetween) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 64)
	if err != nil {
		return err
	}
	if n >= i.Min && n <= i.Max {
		return i.Parser.Parse(d)
	}
	return OutsideBounds(d[0])
}

// Uint is a uint that implements Parser
type Uint struct {
	Data *uint
}

// Parse is an implementation of Parser
func (u Uint) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 0)
	*u.Data = uint(n)
	return err
}

// Uint8 is a uint8 that implements Parser
type Uint8 struct {
	Data *uint8
}

// Parse is an implementation of Parser
func (u Uint8) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 8)
	*u.Data = uint8(n)
	return err
}

// Uint16 is a uint16 that implements Parser
type Uint16 struct {
	Data *uint16
}

// Parse is an implementation of Parser
func (u Uint16) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 16)
	*u.Data = uint16(n)
	return err
}

// Uint32 is a uint32 that implements Parser
type Uint32 struct {
	Data *uint32
}

// Parse is an implementation of Parser
func (u Uint32) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 32)
	*u.Data = uint32(n)
	return err
}

// Uint64 is a uint64 that implements Parser
type Uint64 struct {
	Data *uint64
}

// Parse is an implementation of Parser
func (u Uint64) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 64)
	*u.Data = n
	return err
}

// UintBetween is an implementation of Parser that allows an unsigned integer
// between the Min and Max (inclusively)
type UintBetween struct {
	Parser
	Min, Max uint64
}

// Parse is an implementation of Parser
func (u UintBetween) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 64)
	if err != nil {
		return err
	}
	if n >= u.Min && n <= u.Max {
		return u.Parser.Parse(d)
	}
	return OutsideBounds(d[0])
}

// Float32 is a float32 that implements Parser
type Float32 struct {
	Data *float32
}

// Parse is an implementation of Parser
func (f Float32) Parse(d []string) error {
	n, err := strconv.ParseFloat(d[0], 32)
	*f.Data = float32(n)
	return err
}

// Float64 is a float64 that implements Parser
type Float64 struct {
	Data *float64
}

// Parse is an implementation of Parser
func (f Float64) Parse(d []string) error {
	n, err := strconv.ParseFloat(d[0], 64)
	*f.Data = n
	return err
}

// FloatBetween is an implementation of Parser that allows a float between the
// Min and Max (inclusively)
type FloatBetween struct {
	Parser
	Min, Max float64
}

// Parse is an implementation of Parser
func (f FloatBetween) Parse(d []string) error {
	n, err := strconv.ParseFloat(d[0], 64)
	if err != nil {
		return err
	}
	if n >= f.Min && n <= f.Max {
		return f.Parser.Parse(d)
	}
	return OutsideBounds(d[0])
}

// String is a string that implements Parser
type String struct {
	Data *string
}

// Parse is an implementation of Parser
func (s String) Parse(d []string) error {
	*s.Data = d[0]
	return nil
}

// RegexString is an implementation of Parser that matches a string again the
// given regular expression
type RegexString struct {
	Data  *string
	Regex *regexp.Regexp
}

// Parse is an implementation of Parser
func (r RegexString) Parse(d []string) error {
	if r.Regex.MatchString(d[0]) {
		*r.Data = d[0]
		return nil
	}
	return NoRegexMatch(d[0])
}

// RequiredString is like String except that it returns an error when it gets
// an empty string
type RequiredString String

// Parse is an implementation of Parser
func (r RequiredString) Parse(d []string) error {
	if d[0] == "" {
		return Empty{}
	}
	*r.Data = d[0]
	return nil
}

var formats = [...]string{
	time.ANSIC,
	time.Kitchen,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RubyDate,
	time.Stamp,
	time.StampNano,
	time.StampMicro,
	time.StampMilli,
	time.UnixDate,
	"2/1/06 15:04:05",
	"2/1/2006 15:04:05",
	"02/01/06 15:04:05",
	"02/01/2006 15:04:05",
	"2/1/06 15:04",
	"2/1/2006 15:04",
	"02/01/06 15:04",
	"02/01/2006 15:04",
	"15:04:05 2/1/06",
	"15:04:05 2/1/2006",
	"15:04:05 02/01/06",
	"15:04:05 02/01/2006",
	"15:04 2/1/06",
	"15:04 2/1/2006",
	"15:04 02/01/06",
	"15:04 02/01/2006",
	"2-1-06 15:04:05",
	"2-1-2006 15:04:05",
	"02-01-06 15:04:05",
	"02-01-2006 15:04:05",
	"2-1-06 15:04",
	"2-1-2006 15:04",
	"02-01-06 15:04",
	"02-01-2006 15:04",
	"15:04:05 2-1-06",
	"15:04:05 2-1-2006",
	"15:04:05 02-01-06",
	"15:04:05 02-01-2006",
	"15:04 2-1-06",
	"15:04 2-1-2006",
	"15:04 02-01-06",
	"15:04 02-01-2006",
}

// Time is a time.Time that implements Parser
type Time struct {
	Data *time.Time
}

// Parse is an implementation of Parser
func (t Time) Parse(d []string) error {
	for _, format := range formats {
		pt, err := time.Parse(format, d[0])
		if err == nil {
			*t.Data = pt
			return nil
		}
	}
	return UnknownFormat(d[0])
}

// UnixTime is like time except that it expects the time to be a Unix timestamp
type UnixTime struct {
	Data *time.Time
}

// Parse is an implementation of Parser
func (u UnixTime) Parse(d []string) error {
	i, err := strconv.ParseInt(d[0], 10, 64)
	if err != nil {
		return err
	}
	*u.Data = time.Unix(i, 0)
	return nil
}

// TimeFormat is like a Time except that it allows a specific format to be
// specified
type TimeFormat struct {
	Data   *time.Time
	Format string
}

// Parse is an implementation of Parser
func (t TimeFormat) Parse(d []string) error {
	pt, err := time.Parse(t.Format, d[0])
	if err != nil {
		*t.Data = pt
	}
	return err
}

// Errors

// OutsideBounds is an error returned from a Parser when the parsed value falls
// outside of the range given as an acceptable value
type OutsideBounds string

func (OutsideBounds) Error() string {
	return "value received was outside of specified bounds"
}

// NoRegexMatch is an error returned from RegexString.Parse when the given
// string does not match the given regular expression
type NoRegexMatch string

func (NoRegexMatch) Error() string {
	return "string did not match given regular expression"
}

// Empty is returned when a field contains only an empty string and that is not
// allowed
type Empty struct{}

func (Empty) Error() string {
	return "field is empty"
}

// UnknownFormat is an error returned from Time.Parse when it cannot determine
// the time format of the given string
type UnknownFormat string

func (UnknownFormat) Error() string {
	return "unknown time format"
}
