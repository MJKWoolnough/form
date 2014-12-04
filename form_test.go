package form

import (
	"github.com/MJKWoolnough/equaler"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		a, b ParserList
		data string
	}{
		{
			ParserList{
				"testA": new(Int8),
				"testB": new(Int8),
				"testC": new(Int8),
				"testD": new(Uint16),
				"testE": new(Float32),
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

func newInt8(v Int8) *Int8 {
	return &v
}

func newUint16(v Uint16) *Uint16 {
	return &v
}

func newFloat32(v Float32) *Float32 {
	return &v
}

func (i *Int8) Equal(e equaler.Equaler) bool {
	if d, ok := e.(*Int8); ok {
		return *d == *i
	}
	return false
}

func (u *Uint16) Equal(e equaler.Equaler) bool {
	if d, ok := e.(*Uint16); ok {
		return *d == *u
	}
	return false
}

func (f *Float32) Equal(e equaler.Equaler) bool {
	if d, ok := e.(*Float32); ok {
		return *d == *f
	}
	return false
}
