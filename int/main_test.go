package main

import "testing"

func TestSum(t *testing.T) {
	nums := Sum(20, 20)
	expected := 40

	if nums != expected {
		t.Error("Expected:", expected, "Result:", nums)
	}
}
