package blacksmith

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type Validation struct {
	Data   url.Values
	Errors map[string]string
}

func (bls *Blacksmith) Validator(data url.Values) *Validation {
	return &Validation{
		Errors: make(map[string]string),
		Data:   data,
	}
}

func (v *Validation) Valid() bool {
	return len(v.Errors) == 0
}

// AddError() insert a new key and message error into validation.Errors map
func (v *Validation) AddError(key, message string) {
	// if the error key does not exists previously in the Errors map then
	// insert on it the new key and error message
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Has() check if certain field exists inside a form received through the http request
func (v *Validation) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required() iterates on the form fields and checks for each one must not be empty
func (v *Validation) Required(r *http.Request, fields ...string) {
	for _, field := range fields {
		value := r.Form.Get(field)
		if strings.TrimSpace(value) == "" {
			v.AddError(field, "This field is mandatory")
		}
	}

}

func (v *Validation) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

//isEmail use govalidator to validate an email
func (v *Validation) IsEmail(field, value string) {
	if !govalidator.IsEmail(value) {
		v.AddError(field, "Invalid email address")
	}
}

//isEmail use govalidator to validate an email
func (v *Validation) IsInt(field, value string) {
	_, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(field, "mus be an integer")
	}
}

//isEmail use govalidator to validate an email
func (v *Validation) IsFloat(field, value string) {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		v.AddError(field, "must be a floating point number")
	}
}

func (v *Validation) IsDateISO(field, value string) {
	_, err := time.Parse("2006-01-02", value)
	if err != nil {
		v.AddError(field, "wrong date format, must be YYY-MM-DD")
	}
}

func (v *Validation) NoSpaces(field, value string) {
	if !govalidator.HasWhitespace(value) {
		v.AddError(field, "No white spaces please!")
	}
}
