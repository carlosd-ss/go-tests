package main

import "testing"

func TestPer(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Per(rectangle)
	expected := 40.0
	if result != expected {
		t.Error("Expected:", expected, "Result:", result)
	}
}

func TestAre(t *testing.T) {
	testare := []struct {
		form     Form
		expected float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range testare {
		result := tt.form.Area()
		if result != tt.expected {
			t.Error("Expected:", tt.expected, "Result:", result)
		}
	}

}
