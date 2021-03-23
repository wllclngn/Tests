package greetings

import (
	"errors"
	"regexp"
	"testing"
)

// https://golang.org/doc/tutorial/add-a-test

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func Hello(x string) (string, error) {
	if x == "" {
		return "", errors.New("empty name")
	}

	return x, nil
}
