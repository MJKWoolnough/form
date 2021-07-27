package form

import (
	"math"
	"reflect"
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
	} {
		output := createTypeMap(test.Input)
		if !reflect.DeepEqual(output, test.Output) {
			t.Errorf("test %d: expecting output %v, got %v", n+1, test.Output, output)
		}
	}
}
