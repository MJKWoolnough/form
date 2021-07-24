package form

import (
	"reflect"
	"strconv"
)

type processor interface {
	process(reflect.Value, []string) error
}

type inum struct {
	min, max int64
	bits     int
}

func newInum(tags reflect.StructTag, bits int) inum {
	i := inum{
		min:  0x8000000000000000,
		max:  0x7FFFFFFFFFFFFFFF,
		bits: bits,
	}
	if m := tags.Get("min"); m != "" {
		im, err := strconv.ParseInt(m, 10, bits)
		if err == nil {
			i.min = im
		}
	}
	if m := tags.Get("max"); m != "" {
		im, err := strconv.ParseInt(m, 10, bits)
		if err == nil {
			i.max = im
		}
	}
	return i
}

func (i inum) process(v reflect.Value, data []string) error {
	num, err := strconv.ParseInt(data[0], 10, i.bits)
	if err != nil {
		return err
	}
	if num < i.min || num > i.max {
		return ErrNotInRange
	}
	v.SetInt(num)
	return nil
}
