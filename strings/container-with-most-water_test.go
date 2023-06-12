package main

import "testing"

func TestMaxWaterArea(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	actual := MaxWaterArea(input)
	expected := 49

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = []int{1, 1}
	actual = MaxWaterArea(input)
	expected = 1

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = []int{2, 3, 10, 5, 7, 8, 9}
	actual = MaxWaterArea(input)
	expected = 36

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}
}
