package main

import (
	"reflect"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	actual := TwoSumSorted(nums, target)
	expected := []int{1, 2}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\" and \"%v\": expected result to be %v, got %v", nums, target, expected, actual)
	}

	nums = []int{2, 3, 4}
	target = 6
	actual = TwoSumSorted(nums, target)
	expected = []int{1, 3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\" and \"%v\": expected result to be %v, got %v", nums, target, expected, actual)
	}

	nums = []int{-1, 0}
	target = -1
	actual = TwoSumSorted(nums, target)
	expected = []int{1, 2}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\" and \"%v\": expected result to be %v, got %v", nums, target, expected, actual)
	}
}
