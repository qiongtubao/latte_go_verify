package latte_go_verify

import (
	"testing"
)

type Student struct {
	name string `minLen:"4" maxLen:"8"`
}

func Test_ObjectVerify(t *testing.T) {
	student := Student{"hello"}
	verify := Verify{student}
	err := verify.Set("name", "wor")
	if err != nil {
		t.Error(err)
	}
}
