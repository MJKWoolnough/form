package form

import (
	"strconv"
	"strings"
)

type Bool bool

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

type Int int

func (i *Int) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 0)
	*i = Int(n)
	return err
}

type Int8 int8

func (i *Int8) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 8)
	*i = Int8(n)
	return err
}

type Int16 int16

func (i *Int16) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 16)
	*i = Int16(n)
	return err
}

type Int32 int32

func (i *Int32) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 32)
	*i = Int32(n)
	return err
}

type Int64 int64

func (i *Int64) Parse(d []string) error {
	n, err := strconv.ParseInt(d[0], 10, 64)
	*i = Int64(n)
	return err
}

type Uint uint

func (u *Uint) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 0)
	*u = Uint(n)
	return err
}

type Uint8 uint8

func (u *Uint8) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 8)
	*u = Uint8(n)
	return err
}

type Uint16 uint16

func (u *Uint16) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 16)
	*u = Uint16(n)
	return err
}

type Uint32 uint32

func (u *Uint32) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 32)
	*u = Uint32(n)
	return err
}

type Uint64 uint64

func (u *Uint64) Parse(d []string) error {
	n, err := strconv.ParseUint(d[0], 10, 64)
	*u = Uint64(n)
	return err
}

type Float32 float32

func (f *Float32) Parse(d []string) error {
	n, err := strconv.ParseFloat(d[0], 32)
	*f = Float32(n)
	return err
}

type Float64 float64

func (f *Float64) Parse(d []string) error {
	n, err := strconv.ParseFloat(d[0], 64)
	*f = Float64(n)
	return err
}

type String string

func (s *String) Parse(d []string) error {
	*s = String(d[0])
	return nil
}
