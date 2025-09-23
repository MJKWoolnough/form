package form_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"vimagination.zapto.org/form"
)

func Example() {
	type Form struct {
		Username string   `form:"user,required,regex=^[a-zA-Z0-9_]{3,16}$"`
		Age      int      `form:"age,required,min=18,max=120"`
		Active   bool     `form:"isActive"`
		Tags     []string `form:"tags,regex=^[a-z]+$"`
		IQ       *float64 `form:"iq,min=0,max=200"`
		Token    string   `form:"token,post,required"`
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		var f Form

		if err := form.Process(r, &f); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		json.NewEncoder(w).Encode(f)
	}

	data := url.Values{}
	data.Set("user", "exampleUser")
	data.Set("age", "30")
	data.Set("isActive", "true")
	data.Add("tags", "example")
	data.Add("tags", "tag")
	data.Set("iq", "123.45")
	data.Set("token", "my_token")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	handler(w, req)

	fmt.Println(w.Body.String())

	// Output: {"Username":"exampleUser","Age":30,"Active":true,"Tags":["example","tag"],"IQ":123.45,"Token":"my_token"}
}
