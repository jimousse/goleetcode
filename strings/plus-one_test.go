package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	input := []int{1, 2, 3, 9}
	actual := PlusOne(input)
	expected := []int{1, 2, 4, 0}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = []int{9, 9, 9, 9}
	actual = PlusOne(input)
	expected = []int{1, 0, 0, 0, 0}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = []int{1, 0}
	actual = PlusOne(input)
	expected = []int{1, 1}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}
}
