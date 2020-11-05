package main

import "testing"

func Test_hellow(t *testing.T) {
	result := Hellow("Carlos")
	expected := "Hellow Carlos"

	if result != expected {
		t.Errorf("Result '%s' expected '%s'", result, expected)

	}

}
