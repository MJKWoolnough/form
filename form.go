// Package form provides an easy to use way to parse form values from an HTTP
// request into a struct or other data object
package form // import "vimagination.zapto.org/form"

import (
	"net/http"
	"reflect"
	"sync"
)

type processorDetails struct {
	processor
	Post, Required bool
	Index          []int
}

type typeMap map[string]processorDetails

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
	tm, ok := typeMaps[t]
	if ok {
		return tm
	}
	tm = make(typeMap)
	for i := 0; i < t.Len(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		name := f.Name
		if n := f.Tag.Get("form"); n == "-" {
			continue
		} else if n != "" {
			name = n
		}
		switch f.Type.Kind() {
		case reflect.Int8:
			tm[name] = newInum(f.Type, 8)
		case reflect.Int16:
			tm[name] = newInum(f.Type, 8)
		case reflect.Int32:
			tm[name] = newInum(f.Type, 8)
		case reflect.Int64:
			tm[name] = newInum(f.Type, 8)
		case reflect.Int:
			tm[name] = newInum(f.Type, 8)
		}
	}
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
	var errors Errors
	for key, pd := range tm {
		var (
			val []string
			ok  bool
		)
		if pd.Post {
			val, ok = r.PostForm[key]
		} else {
			val, ok = r.Form[key]
		}
		if ok {
			if err := pd.processor.process(v.FieldByIndex(pd.Index), val); err != nil {
				errors = append(errors, ErrProcessingFailed{
					Key:   key,
					Error: err,
				})
			}
		} else if pd.Required {
			errors = append(errors, ErrRequiredMissing(key))
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}
