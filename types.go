package form

import (
	"strconv"
	"strings"
)

// Bool is a bool that implements Parser
type Bool bool

// Parse is an implementation of Parser
func (b *Bool) Parse(d []string) error {
	switch strings.ToLower(d[0]) {
	case "on", "yes", "y":
		*b = true
	case "off", "no", "n":
		*b = false
	default:
		c, err := strconv.ParseBool(d[0])
		*b = Bool(c)
		return err
	}
	return nil
}

// Int is an int that implements Parser
type Int int

// Parse is an implementation of Parser
func (i *Int) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 0)
	*i = Int(n)
	return err
}

// Int8 is an int8 that implements Parser
type Int8 int8

// Parse is an implementation of Parser
func (i *Int8) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 8)
	*i = Int8(n)
	return err
}

// Int16 is an int16 that implements Parser
type Int16 int16

// Parse is an implementation of Parser
func (i *Int16) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 16)
	*i = Int16(n)
	return err
}

// Int32 is an int32 that implements Parser
type Int32 int32

// Parse is an implementation of Parser
func (i *Int32) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 32)
	*i = Int32(n)
	return err
}

// Int64 is an int64 that implements Parser
type Int64 int64

// Parse is an implementation of Parser
func (i *Int64) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 64)
	*i = Int64(n)
	return err
}

// Uint is a uint that implements Parser
type Uint uint

// Parse is an implementation of Parser
func (u *Uint) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 0)
	*u = Uint(n)
	return err
}

// Uint8 is a uint8 that implements Parser
type Uint8 uint8

// Parse is an implementation of Parser
func (u *Uint8) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 8)
	*u = Uint8(n)
	return err
}

// Uint16 is a uint16 that implements Parser
type Uint16 uint16

// Parse is an implementation of Parser
func (u *Uint16) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 16)
	*u = Uint16(n)
	return err
}

// Uint32 is a uint32 that implements Parser
type Uint32 uint32

// Parse is an implementation of Parser
func (u *Uint32) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 32)
	*u = Uint32(n)
	return err
}

// Uint64 is a uint64 that implements Parser
type Uint64 uint64

// Parse is an implementation of Parser
func (u *Uint64) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 64)
	*u = Uint64(n)
	return err
}

// Float32 is a float32 that implements Parser
type Float32 float32

// Parse is an implementation of Parser
func (f *Float32) Parse(d []string) error {
	n, err := strconv.ParseFloat(d[0], 32)
	*f = Float32(n)
	return err
}

// Float64 is a float64 that implements Parser
type Float64 float64

// Parse is an implementation of Parser
func (f *Float64) Parse(d []string) error {
	n, err := strconv.ParseFloat(d[0], 64)
	*f = Float64(n)
	return err
}

// String is a string that implements Parser
type String string

// Parse is an implementation of Parser
func (s *String) Parse(d []string) error {
	*s = String(d[0])
	return nil
}
