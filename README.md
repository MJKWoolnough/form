# form

[![CI](https://github.com/MJKWoolnough/form/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/form/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/form.svg)](https://pkg.go.dev/vimagination.zapto.org/form)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/form)](https://goreportcard.com/report/vimagination.zapto.org/form)

--
    import "vimagination.zapto.org/form"

Package form provides an easy to use way to parse form values from an HTTP request into a struct.

## Highlights

 - Parse form values directly into Go structs.
 - Use struct tags to define limitations and requirements for the field.
 - Supports nested and anonymous structs, slices, and pointers.
 - Processors for all basic types and allows custom processing by including a `ParseForm([]string) error` method on a type.

## Usage

```go
package main

import (
	"encoding/json"
	"net/http"

	"vimagination.zapto.org/form"
)

type Form struct {
	Username string   `form:"user,required,regex=^[a-zA-Z0-9_]{3,16}$"`
	Age      int      `form:"age,required,min=18,max=120"`
	Active   bool     `form:"isActive"`
	Tags     []string `form:"tags,regex=^[a-z]+$"`
	IQ       *float64 `form:"iq,min=0,max=200"`
	Token    string   `form:"token,post,required"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var f Form

	if err := form.Process(r, &f); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	json.NewEncoder(w).Encode(form)
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/form
