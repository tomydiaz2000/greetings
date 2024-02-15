package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) { // the object t is used for reporting and logging
	name := "tom"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("tom")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("tom") = %q, %v, want %q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want %q, error`, message, err, "")
	}
}