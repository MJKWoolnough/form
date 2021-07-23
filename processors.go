package form

import "reflect"

type processor interface {
	process(reflect.Value, []string) error
}
