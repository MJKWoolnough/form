// Package form provides an easy to use way to parse form values from an HTTP
// request into a struct or other data object
package form // import "vimagination.zapto.org/form"

import (
	"errors"
	"net/http"
	"reflect"
)

type processor struct {
	Post, Required bool
	Index          []int
}

func (p processor) process(v reflect.Value, data []string) error {
	return nil
}

type typeMap map[string]processor

func createTypeMap(t reflect.Type) typeMap {
	return nil
}

func ProcessForm(r *http.Request, fv interface{}) error {
	v := reflect.ValueOf(fv)
	if v.Kind() != reflect.Ptr {
		return ErrNeedPointer
	}
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ErrNeedStruct
	}
	tm := createTypeMap(v.Type())
	if err := r.ParseForm(); err != nil {
		return err
	}
	for key, processor := range tm {
		var (
			val []string
			ok  bool
		)
		if processor.Post {
			val, ok = r.PostForm[key]
		} else {
			val, ok = r.Form[key]
		}
		if ok {
			if err := processor.process(v.FieldByIndex(processor.Index), val); err != nil {

			}
		} else if processor.Required {

		}
	}
	return nil
}

var (
	ErrNeedPointer = errors.New("need pointer to type")
	ErrNeedStruct  = errors.New("need struct type")
)
