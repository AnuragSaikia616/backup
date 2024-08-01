package main

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	NonFieldErrors []string
	fromErrors     map[string]string
}

func (v *Validator) AddField(key, value string) {
	if v.fromErrors == nil {
		v.fromErrors = make(map[string]string)
	}

	if _, exist := v.fromErrors[key]; !exist {
		v.fromErrors[key] = value

	}

}

func (v *Validator) valid() bool {
	return len(v.fromErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddField(key, message)
	}
}

func (v *Validator) isBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func (v *Validator) maxChars(nchars int, value string) bool {
	return utf8.RuneCountInString(value) <= nchars
}

func (v *Validator) minChars(nchars int, value string) bool {
	return utf8.RuneCountInString(value) >= nchars
}

func (v *Validator) matches(rx *regexp.Regexp, value string) bool {
	return rx.MatchString(value)
}

func (v *Validator) permitted(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func (v *Validator) Error(w http.ResponseWriter) {

}
