package main

import "testing"

func TestFor(t *testing.T) {
	repetitions := Rep("a")
	expected := "aaaaa"
	if repetitions != expected {
		t.Error("Expected:", expected, "Result:", repetitions)
	}
}

func BenchmarkFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rep("a")
	}
}
