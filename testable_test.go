package main

import "testing"

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	r := add(1, 2)
	if r != 3 {
		t.Error("expect", 3, "actual", r)
	}
}
