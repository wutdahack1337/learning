package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Minh"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Minh")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Minh") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string, checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
