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
		if im, err := strconv.ParseInt(m, 10, bits); err == nil {
			i.min = im
		}
	}

	if m := tags.Get("max"); m != "" {
		if im, err := strconv.ParseInt(m, 10, bits); err == nil {
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
		if um, err := strconv.ParseUint(m, 10, bits); err == nil {
			u.min = um
		}
	}

	if m := tags.Get("max"); m != "" {
		if um, err := strconv.ParseUint(m, 10, bits); err == nil {
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
		if um, err := strconv.ParseFloat(m, bits); err == nil {
			f.min = um
		}
	}

	if m := tags.Get("max"); m != "" {
		if um, err := strconv.ParseFloat(m, bits); err == nil {
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
	if s.regex != nil && !s.regex.MatchString(data[0]) {
		return ErrNoMatch
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

type boolChars [256]uint8

type boolResult uint8

const (
	resultNone boolResult = iota
	resultTrue
	resultFalse
)

type boolState struct {
	boolChars
	result boolResult
}

var bools = [...]boolState{
	{},
	{ // 1
		boolChars: boolChars{
			'0': 2,
			'1', 3,
			'F': 4,
			'f': 4,
			'N': 8,
			'n': 8,
			'O': 9,
			'o': 9,
			'T': 11,
			't': 11,
			'Y': 14,
			'y': 14,
		},
	},
	{ // 2: '0', 'false', 'off'
		result: resultFalse,
	},
	{ // 3: '1', 'on', 'true', 'yes'
		result: resultTrue,
	},
	{ // 4: 'f',
		boolChars: boolChars{
			'A': 5,
			'a': 5,
		},
		result: resultFalse,
	},
	{ // 5: 'fa',
		boolChars: boolChars{
			'L': 6,
			'l': 6,
		},
	},
	{ // 6: 'fal',
		boolChars: boolChars{
			'S': 7,
			's': 7,
		},
	},
	{ // 7: 'fals',
		boolChars: boolChars{
			'E': 2,
			'e': 2,
		},
	},
	{ // 8: 'n',
		boolChars: boolChars{
			'O': 2,
			'o': 2,
		},
		result: resultFalse,
	},
	{ // 9: 'o',
		boolChars: boolChars{
			'F': 10,
			'f': 10,
			'N': 3,
			'n': 3,
		},
	},
	{ // 10: 'of',
		boolChars: boolChars{
			'F': 2,
			'f': 2,
		},
	},
	{ // 11: 't',
		boolChars: boolChars{
			'R': 12,
			'r': 12,
		},
		result: resultTrue,
	},
	{ // 12: 'tr',
		boolChars: boolChars{
			'U': 13,
			'u': 13,
		},
	},
	{ // 13: 'tru',
		boolChars: boolChars{
			'E': 3,
			'e': 3,
		},
	},
	{ // 14: 'y',
		boolChars: boolChars{
			'E': 15,
			'e': 15,
		},
		result: resultTrue,
	},
	{ // 15: 'ye',
		boolChars: boolChars{
			'S': 3,
			's': 3,
		},
	},
}

func (boolean) process(v reflect.Value, data []string) error {
	pos := uint8(1)

	for n := range data[0] {
		if pos = bools[pos].boolChars[data[0][n]]; pos == 0 {
			break
		}
	}

	result := bools[pos].result

	if result == resultNone {
		return ErrInvalidBoolean
	}

	v.SetBool(result == resultTrue)

	return nil
}

type slice struct {
	processor
	typ reflect.Type
}

func (s slice) process(v reflect.Value, data []string) error {
	if v.Cap() >= len(data) {
		v.SetLen(len(data))
	} else {
		v.Set(reflect.MakeSlice(s.typ, len(data), len(data)))
	}

	var errs Errors

	for n := range data {
		if err := s.processor.process(v.Index(n), data[n:]); err != nil {
			if errs == nil {
				errs = make(Errors, len(data))
			}

			errs[n] = err
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
