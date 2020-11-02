package main

import "testing"

func Test_hellow(t *testing.T) {
	result := hellow("Carlos")
	expected := "Hellow Carlos"

	if result != expected {
		t.Errorf("Result '%s' expected '%s'", result, expected)

	}

}
