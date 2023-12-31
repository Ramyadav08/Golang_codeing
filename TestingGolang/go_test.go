package main

import (
	"testing"
)

// test function
func TestReturnGeeks(t *testing.T) {
	actualString := ReturnGeeks()
	expectedString := "geeks"
	if actualString != expectedString{
		t.Errorf("Expected String(%s) is not same as"+
		" actual string (%s)", expectedString,actualString)
	}
}

// function to Benchmark ReturnGeeks()
func BenchmarkGeeks(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ReturnGeeks()
    }
}