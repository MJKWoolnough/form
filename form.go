package form

import (
	"net/url"
)

type Errors map[string]error

func (Errors) Error() string {
	return "errors encountered"
}

type Parser interface {
	Parse([]string) error
}

type ParserList map[string]Parser

func (p ParserList) ParserList() ParserList {
	return p
}

type parserLister interface {
	ParserList() ParserList
}

func ParseValue(name string, value Parser, data url.Values) error {
	return Parse(ParserList{name: value}, data)
}

func Parse(p parserLister, data url.Values) error {
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
