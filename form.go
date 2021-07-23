// Package form provides an easy to use way to parse form values from an HTTP
// request into a struct or other data object
package form // import "vimagination.zapto.org/form"

import (
	"errors"
	"net/http"
	"reflect"
	"sync"
)

type processor struct {
	Post, Required bool
	Index          []int
}

func (p processor) process(v reflect.Value, data []string) error {
	return nil
}

type typeMap map[string]processor

var (
	tmMu     sync.RWMutex
	typeMaps = make(map[reflect.Type]typeMap)
)

func getTypeMap(t reflect.Type) typeMap {
	tmMu.RLock()
	tm, ok := typeMaps[t]
	tmMu.RUnlock()
	if ok {
		return tm
	}
	tmMu.Lock()
	tm = createTypeMap(t)
	tmMu.Unlock()
	return tm
}

func createTypeMap(t reflect.Type) typeMap {
	tm, ok = typeMaps[t]
	if ok {
		return tm
	}
	tm = make(typeMap)
	typeMaps[t] = tm
	return tm
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
	tm := getTypeMap(v.Type())
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
