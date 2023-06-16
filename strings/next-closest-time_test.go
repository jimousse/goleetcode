package main

import (
	"testing"
)

func TestNextClosestTime(t *testing.T) {
	input := "19:34"
	actual := NextClosestTime(input)
	expected := "19:39"

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = "23:59"
	actual = NextClosestTime(input)
	expected = "22:22"

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = "01:34"
	actual = NextClosestTime(input)
	expected = "01:40"

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

}
