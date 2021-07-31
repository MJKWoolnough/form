package form

import (
	"math"
	"reflect"
	"regexp"
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
		min:  math.MinInt64,
		max:  math.MaxInt64,
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

type unum struct {
	min, max uint64
	bits     int
}

func newUnum(tags reflect.StructTag, bits int) unum {
	u := unum{
		max:  math.MaxUint64,
		bits: bits,
	}
	if m := tags.Get("min"); m != "" {
		um, err := strconv.ParseUint(m, 10, bits)
		if err == nil {
			u.min = um
		}
	}
	if m := tags.Get("max"); m != "" {
		um, err := strconv.ParseUint(m, 10, bits)
		if err == nil {
			u.max = um
		}
	}
	return u
}

func (u unum) process(v reflect.Value, data []string) error {
	num, err := strconv.ParseUint(data[0], 10, u.bits)
	if err != nil {
		return err
	}
	if num < u.min || num > u.max {
		return ErrNotInRange
	}
	v.SetUint(num)
	return nil
}

type float struct {
	min, max float64
	bits     int
}

func newFloat(tags reflect.StructTag, bits int) float {
	f := float{
		min:  -math.MaxFloat64,
		max:  math.MaxFloat64,
		bits: bits,
	}
	if m := tags.Get("min"); m != "" {
		um, err := strconv.ParseFloat(m, bits)
		if err == nil {
			f.min = um
		}
	}
	if m := tags.Get("max"); m != "" {
		um, err := strconv.ParseFloat(m, bits)
		if err == nil {
			f.max = um
		}
	}
	return f
}

func (f float) process(v reflect.Value, data []string) error {
	num, err := strconv.ParseFloat(data[0], f.bits)
	if err != nil {
		return err
	}
	if num < f.min || num > f.max {
		return ErrNotInRange
	}
	v.SetFloat(num)
	return nil
}

type str struct {
	regex *regexp.Regexp
}

func newString(tags reflect.StructTag) str {
	if r := tags.Get("regex"); r != "" {
		if re, err := regexp.Compile(r); err == nil {
			return str{
				regex: re,
			}
		}
	}
	return str{}
}

func (s str) process(v reflect.Value, data []string) error {
	if s.regex != nil {
		if !s.regex.MatchString(data[0]) {
			return ErrNoMatch
		}
	}
	v.SetString(data[0])
	return nil
}

type boolean struct{}

func matchString(a string, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for n, c := range b {
		if a[n]|32 != c {
			return false
		}
	}
	return true
}

var (
	trues = [...][]byte{
		{'1'},
		{'y'},
		{'t'},
		{'o', 'n'},
		{'y', 'e', 's'},
		{'t', 'r', 'u', 'e'},
	}
	falses = [...][]byte{
		{'0'},
		{'n'},
		{'g'},
		{'o', 'f', 'f'},
		{'n', 'o'},
		{'f', 'a', 'l', 's', 'e'},
	}
)

func (boolean) process(v reflect.Value, data []string) error {
	for _, b := range trues {
		if matchString(data[0], b) {
			v.SetBool(true)
			return nil
		}
	}
	for _, b := range falses {
		if matchString(data[0], b) {
			v.SetBool(false)
			return nil
		}
	}
	return ErrInvalidBoolean
}

type slice struct {
	processor
	typ reflect.Type
}

func (s slice) process(v reflect.Value, data []string) error {
	v.Set(reflect.MakeSlice(s.typ, 1, len(data)))
	l := 1
	errs := make(Errors, 0)
	for n := range data {
		if err := s.processor.process(v.Index(l-1), data[n:]); err != nil {
			errs = append(errs, err)
		} else if l < len(data) {
			l++
			v.SetLen(l)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

type pointer struct {
	processor
	typ reflect.Type
}

func (p pointer) process(v reflect.Value, data []string) error {
	pv := reflect.New(p.typ)
	if err := p.processor.process(pv.Elem(), data); err != nil {
		return err
	}
	v.Set(pv)
	return nil
}

type formParser interface {
	ParseForm([]string) error
}

type inter bool

func (i inter) process(v reflect.Value, data []string) error {
	if i {
		v = v.Addr()
	}
	return v.Interface().(formParser).ParseForm(data)
}
