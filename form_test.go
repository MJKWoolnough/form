package form

import (
	"net/url"
	"testing"

	"github.com/MJKWoolnough/equaler"
)

func TestParse(t *testing.T) {
	tests := []struct {
		a, b ParserList
		data string
	}{
		{
			ParserList{
				"testA": newInt8(0),
				"testB": newInt8(0),
				"testC": newInt8(0),
				"testD": newUint16(0),
				"testE": newFloat32(0),
			},
			ParserList{
				"testA": newInt8(1),
				"testB": newInt8(123),
				"testC": newInt8(-127),
				"testD": newUint16(1023),
				"testE": newFloat32(-3.14),
			},
			"testA=1&testB=123&testC=-127&testD=1023&testE=-3.14",
		},
	}

	for n, test := range tests {
		q, _ := url.ParseQuery(test.data)
		err := Parse(test.a, q)
		if err != nil {
			t.Errorf("test %d: unexpected error: %q", n, err)
			continue
		}
		for k, v := range test.b {
			if !v.(equaler.Equaler).Equal(test.a[k].(equaler.Equaler)) {
				t.Errorf("test %d: key %q: expecting %v, got %v", n, k, v, test.a[k])
			}
		}
	}
}

func newInt8(v int8) Int8 {
	return Int8{&v}
}

func newUint16(v uint16) Uint16 {
	return Uint16{&v}
}

func newFloat32(v float32) Float32 {
	return Float32{&v}
}

func (i Int8) Equal(e equaler.Equaler) bool {
	if d, ok := e.(Int8); ok {
		return *d.Data == *i.Data
	}
	return false
}

func (u Uint16) Equal(e equaler.Equaler) bool {
	if d, ok := e.(Uint16); ok {
		return *d.Data == *u.Data
	}
	return false
}

func (f Float32) Equal(e equaler.Equaler) bool {
	if d, ok := e.(Float32); ok {
		return *d.Data == *f.Data
	}
	return false
}
