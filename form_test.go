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
		{
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
	} {
		output := createTypeMap(test.Input)
		if !reflect.DeepEqual(output, test.Output) {
			t.Errorf("test %d: expecting output %v, got %v", n+1, test.Output, output)
		}
	}
}
