package form

import (
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

// String is a string that implements Parser
type String struct {
	Data *string
}

// Parse is an implementation of Parser
func (s String) Parse(d []string) error {
	*s.Data = d[0]
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

// UnknownFormat is an error return from Time.Parse when it cannot determine
// the time format of the given string
type UnknownFormat string

func (UnknownFormat) Error() string {
	return "unknown time format"
}
