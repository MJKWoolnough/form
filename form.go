// Package form provides an easy to use way to parse form values from an HTTP
// request into a struct
package form // import "vimagination.zapto.org/form"

import (
	"net/http"
	"reflect"
	"strings"
	"sync"
)

var interType = reflect.TypeOf((*formParser)(nil)).Elem()

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

func basicTypeProcessor(t reflect.Type, tag reflect.StructTag) processor {
	switch t.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return newInum(tag, t.Bits())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return newUnum(tag, t.Bits())
	case reflect.Float32, reflect.Float64:
		return newFloat(tag, t.Bits())
	case reflect.String:
		return str{}
	case reflect.Bool:
		return boolean{}
	}
	return nil
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
		var required, post bool
		if n := f.Tag.Get("form"); n == "-" {
			continue
		} else if n != "" {
			p := strings.IndexByte(n, ',')
			if p >= 0 {
				if p > 0 {
					name = n[:p]
				}
				rest := n[p:]
				required = strings.Contains(rest, ",required,") || strings.HasPrefix(rest, ",required")
				post = strings.Contains(rest, ",post,") || strings.HasPrefix(rest, ",post")
			} else {
				name = n
			}
		}
		var p processor
		if f.Type.Implements(interType) {
			p = inter(false)
		} else if reflect.PtrTo(f.Type).Implements(interType) {
			p = inter(true)
		} else if k := f.Type.Kind(); k == reflect.Slice || k == reflect.Ptr {
			et := f.Type.Elem()
			s := basicTypeProcessor(et, f.Tag)
			if s == nil {
				continue
			}
			if k == reflect.Slice {
				p = slice{
					processor: s,
					typ:       et,
				}
			} else {
				p = pointer{
					processor: s,
					typ:       et,
				}
			}
		} else if k == reflect.Struct && f.Anonymous {
			for n, p := range createTypeMap(f.Type) {
				if _, ok := tm[n]; !ok {
					tm[n] = processorDetails{
						processor: p.processor,
						Required:  p.Required,
						Post:      p.Post,
						Index:     append(append(make([]int, 0, len(p.Index)+1), i), p.Index...),
					}
				}
			}
			continue
		} else {
			p = basicTypeProcessor(f.Type, f.Tag)
			if p == nil {
				continue
			}
		}
		tm[name] = processorDetails{
			processor: p,
			Required:  required,
			Post:      post,
			Index:     []int{i},
		}
	}
	typeMaps[t] = tm
	return tm
}

// Process parses the form data from the request into the passed value, which
// must be a pointer to a struct.
//
// Form keys are assumed to be the field names unless a 'form' tag is provided
// with an alternate name, for example, in the following struct, the int is
// parse with key 'A' and the bool is parsed with key 'C'.
//
// type Example struct {
//	A int
//	B bool `form:"C"`
// }
//
// Two options can be added to the form tag to modify the processing. The
// 'post' option forces the processer to parse a value from the PostForm field
// of the Request, and the 'required' option will have an error thrown if the
// key in not set.
//
// Number types can also have minimums and maximums checked during processing
// by setting the min and max tags accordingly.
//
// Lastly, a custom data processor can be specified by attaching a method to
// the field type with the following specification:
//
// ParseForm([]string) error
func Process(r *http.Request, fv interface{}) error {
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
					Key: key,
					Err: err,
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
