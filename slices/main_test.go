package main

import (
	"testing"
)

func TestSums(t *testing.T) {
	t.Run("Sum Test", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		result := Sums(nums)
		expected := 15
		if result != expected {
			t.Error("Expected:", expected, "Result:", result)
		}
	})

	t.Run("Slice sum Test", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 100}
		result := Sums(nums)
		expected := 110
		if result != expected {
			t.Error("Expected:", expected, "Result:", result)
		}
	})
}
