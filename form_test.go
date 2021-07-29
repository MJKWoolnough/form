package form

import (
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestCreateTypeMap(t *testing.T) {
	for n, test := range [...]struct {
		Input  reflect.Type
		Output typeMap
	}{
		{ // 1
			Input: reflect.TypeOf(struct {
				A int
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 2
			Input: reflect.TypeOf(struct {
				A int `form:"B"`
			}{}),
			Output: typeMap{
				"B": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 3
			Input: reflect.TypeOf(struct {
				A int `form:",required"`
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Required: true,
					Index:    []int{0},
				},
			},
		},
		{ // 4
			Input: reflect.TypeOf(struct {
				A int `form:",post"`
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Post:  true,
					Index: []int{0},
				},
			},
		},
		{ // 5
			Input: reflect.TypeOf(struct {
				A int `form:",post,required"`
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Post:     true,
					Required: true,
					Index:    []int{0},
				},
			},
		},
		{ // 6
			Input: reflect.TypeOf(struct {
				A int `form:",required,post"`
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Post:     true,
					Required: true,
					Index:    []int{0},
				},
			},
		},
		{ // 7
			Input: reflect.TypeOf(struct {
				A int `min:"-20"`
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  -20,
						max:  math.MaxInt64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 8
			Input: reflect.TypeOf(struct {
				A int `max:"-20"`
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  -20,
					},
					Index: []int{0},
				},
			},
		},
		{ // 9
			Input: reflect.TypeOf(struct {
				A int8
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 8,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 10
			Input: reflect.TypeOf(struct {
				A int16
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 16,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 11
			Input: reflect.TypeOf(struct {
				A int32
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 32,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 12
			Input: reflect.TypeOf(struct {
				A int64
			}{}),
			Output: typeMap{
				"A": {
					processor: inum{
						bits: 64,
						min:  math.MinInt64,
						max:  math.MaxInt64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 13
			Input: reflect.TypeOf(struct {
				A uint
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 64,
						max:  math.MaxUint64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 14
			Input: reflect.TypeOf(struct {
				A uint `min:"10"`
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 64,
						min:  10,
						max:  math.MaxUint64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 15
			Input: reflect.TypeOf(struct {
				A uint `min:"-1"`
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 64,
						max:  math.MaxUint64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 16
			Input: reflect.TypeOf(struct {
				A uint `max:"100"`
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 64,
						max:  100,
					},
					Index: []int{0},
				},
			},
		},
		{ // 17
			Input: reflect.TypeOf(struct {
				A uint8
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 8,
						max:  math.MaxUint64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 18
			Input: reflect.TypeOf(struct {
				A uint16
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 16,
						max:  math.MaxUint64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 19
			Input: reflect.TypeOf(struct {
				A uint32
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 32,
						max:  math.MaxUint64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 20
			Input: reflect.TypeOf(struct {
				A uint
			}{}),
			Output: typeMap{
				"A": {
					processor: unum{
						bits: 64,
						max:  math.MaxUint64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 21
			Input: reflect.TypeOf(struct {
				A float32
			}{}),
			Output: typeMap{
				"A": {
					processor: float{
						bits: 32,
						min:  -math.MaxFloat64,
						max:  math.MaxFloat64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 22
			Input: reflect.TypeOf(struct {
				A float64
			}{}),
			Output: typeMap{
				"A": {
					processor: float{
						bits: 64,
						min:  -math.MaxFloat64,
						max:  math.MaxFloat64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 23
			Input: reflect.TypeOf(struct {
				A float64 `min:"-10"`
			}{}),
			Output: typeMap{
				"A": {
					processor: float{
						bits: 64,
						min:  -10,
						max:  math.MaxFloat64,
					},
					Index: []int{0},
				},
			},
		},
		{ // 24
			Input: reflect.TypeOf(struct {
				A float64 `max:"10"`
			}{}),
			Output: typeMap{
				"A": {
					processor: float{
						bits: 64,
						min:  -math.MaxFloat64,
						max:  10,
					},
					Index: []int{0},
				},
			},
		},
		{ // 25
			Input: reflect.TypeOf(struct {
				A string
			}{}),
			Output: typeMap{
				"A": {
					processor: str{},
					Index:     []int{0},
				},
			},
		},
		{ // 26
			Input: reflect.TypeOf(struct {
				A string `regex:"/aaa/"`
			}{}),
			Output: typeMap{
				"A": {
					processor: str{
						regex: regexp.MustCompile("/aaa/"),
					},
					Index: []int{0},
				},
			},
		},
		{ // 27
			Input: reflect.TypeOf(struct {
				A bool
			}{}),
			Output: typeMap{
				"A": {
					processor: boolean{},
					Index:     []int{0},
				},
			},
		},
		{ // 28
			Input: reflect.TypeOf(struct {
				A []bool
			}{}),
			Output: typeMap{
				"A": {
					processor: slice{
						processor: boolean{},
						typ:       reflect.TypeOf(false),
					},
					Index: []int{0},
				},
			},
		},
	} {
		output := createTypeMap(test.Input)
		if !reflect.DeepEqual(output, test.Output) {
			t.Errorf("test %d: expecting output %v, got %v", n+1, test.Output, output)
		}
	}
}

func TestProcess(t *testing.T) {
	for n, test := range [...]struct {
		Get, Post url.Values
		Output    interface{}
		Err       error
	}{} {
		r := http.Request{
			Method: http.MethodPost,
			URL: &url.URL{
				RawQuery: test.Get.Encode(),
			},
			Header: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			Body: ioutil.NopCloser(strings.NewReader(test.Post.Encode())),
		}
		output := reflect.New(reflect.TypeOf(test.Output))
		err := Process(&r, output.Interface())
		if err != nil {
			if test.Err == nil {
				t.Errorf("test %d: unexpected error: %s", n+1, err)
			} else if err != test.Err {
				t.Errorf("test %d: expecting error: %s\ngot: %s", n+1, test.Err, err)
			}
		} else if test.Err != nil {
			t.Errorf("test %d: got no error when expecting: %s", n+1, test.Err)
		} else if o := output.Elem().Interface(); !reflect.DeepEqual(o, test.Output) {
			t.Errorf("test %d: expecting output: %#v\ngot: %#v", n+1, test.Output, o)
		}
	}
}
