// Package form provides an easy to use way to parse form values from an HTTP
// request into a struct or other data object
package form // import "vimagination.zapto.org/form"

import (
	"net/url"
)

// Errors is an error type that is a map of other errors
type Errors map[string]error

func (Errors) Error() string {
	return "errors encountered"
}

// Parser is an interface used to to parse a specific type
type Parser interface {
	Parse([]string) error
}

// ParserList is a simple implementation of a parserLister that simply returns
// itself
type ParserList map[string]Parser

// ParserList is an implementation of parserLister
func (p ParserList) ParserList() ParserList {
	return p
}

// ParserLister is the main interface for this package. The single method
// ParserList returns a ParserList map of field names to Parser's
type ParserLister interface {
	ParserList() ParserList
}

// ParseValue parses a single values
func ParseValue(name string, value Parser, data url.Values) error {
	return Parse(ParserList{name: value}, data)
}

// Parse parses the given url.Values into the type given
func Parse(p ParserLister, data url.Values) error {
	errs := make(Errors)
	for k, v := range p.ParserList() {
		if d, ok := data[k]; ok {
			if err := v.Parse(d); err != nil {
				errs[k] = err
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
