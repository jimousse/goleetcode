package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	actual := ThreeSum(input)
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = []int{-2, 0, 0, 2, 2}
	actual = ThreeSum(input)
	expected = [][]int{{-2, 0, 2}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	// input = []int{0, 0, 0}
	// actual = ThreeSum(input)
	// expected = [][]int{{0, 0, 0}}

	// if !reflect.DeepEqual(actual, expected) {
	// 	t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	// }
}
