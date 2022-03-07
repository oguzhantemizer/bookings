package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//It creates a custom form structs, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

//New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Checks for required field
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannnot be blank")
		}
	}
}

//Has checks if form field in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == " " {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

//Valid returns true if there ara no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//Minlength checks for miniimum characters long
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must me at least %d characters long", length))
		return false
	}
	return true
}

//Checks for valid email adress
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email adress")
	}
}
