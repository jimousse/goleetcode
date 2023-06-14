package main

import "testing"

func TestCanJump(t *testing.T) {
	input := []int{2, 3, 1, 1, 4}
	actual := CanJump(input)
	expected := true

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = []int{3, 2, 1, 0, 4}
	actual = CanJump(input)
	expected = false

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = []int{0}
	actual = CanJump(input)
	expected = true

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}
}
