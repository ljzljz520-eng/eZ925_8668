package model

import "testing"

func TestValidation(t *testing.T) {
	p := Profile{ID: "p", Name: "n", AccessCode: "1234"}
	if !p.Valid() {
		t.Fatal()
	}
	if !IsSafePath("x/y") {
		t.Fatal()
	}
}
