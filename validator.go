package blacksmith

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

// Validation type to store form data and validations errors
type Validation struct {
	Data   url.Values
	Errors map[string]string
}

// Validator() receiver func to initialize the custom validator type
func (bls *Blacksmith) Validator(data url.Values) *Validation {
	return &Validation{
		Errors: make(map[string]string),
		Data:   data,
	}
}

// Valid() returns true if there is no error un Validation.Errors map
func (v *Validation) Valid() bool {
	return len(v.Errors) == 0
}

// AddError() insert a new key and message error into validation.Errors map
func (v *Validation) AddError(key, message string) {
	// if the error key does not exists previously in the Errors map then
	// insert in the v.Errors map the new key and error message
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Has() check if certain field exists inside a form received through the http request
//this func do not use govalidator
func (v *Validation) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required() iterates on the form fields and check for that each one must not be empty
func (v *Validation) Required(r *http.Request, fields ...string) {
	for _, field := range fields {
		value := r.Form.Get(field)
		if strings.TrimSpace(value) == "" {
			v.AddError(field, "This field is mandatory")
		}
	}

}

// Check() use a simple bool to add or not en error to Validation.Errors map
// will be used t 	 	o
func (v *Validation) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

//isEmail use govalidator to validate an email in a form field
func (v *Validation) IsEmail(field, value string) {
	if !govalidator.IsEmail(value) {
		v.AddError(field, "Invalid email address")
	}
}

//IsInt use govalidator to validate an int value in a form field
func (v *Validation) IsInt(field, value string) {
	_, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(field, "must be an integer")
	}
}

//isFloeat use govalidator to validate a float value in a form field
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

// NoSpaces check for a field with white spaces between the words
func (v *Validation) NoSpaces(field, value string) {
	if !govalidator.HasWhitespace(value) {
		v.AddError(field, "No white spaces please!")
	}
}
