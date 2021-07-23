package form

import (
	"errors"
	"reflect"
	"strconv"
)

type processor interface {
	process(reflect.Value, []string) error
}

type inum8 struct {
	min, max int8
}

func newInum8(tags reflect.StructTag) inum8 {
	i := inum8{
		min: -128,
		max: 127,
	}
	if m := tags.Get("min"); m != "" {
		im, err := strconv.ParseInt(m, 10, 8)
		if err == nil {
			i.min = int8(im)
		}
	}
	if m := tags.Get("max"); m != "" {
		im, err := strconv.ParseInt(m, 10, 8)
		if err == nil {
			i.max = int8(im)
		}
	}
	return i
}

func (i inum8) process(v reflect.Value, data []string) error {
	num, err := strconv.ParseInt(data[0], 10, 8)
	if err != nil {
		return err
	}
	if num < i.min || num > i.max {
		return ErrNotInRange
	}
	v.SetInt(num)
	return nil
}

// Errors
var (
	ErrNotInRange = errors.New("value not in valid range")
)
