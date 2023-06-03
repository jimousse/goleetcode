package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	actual := TwoSum(nums, target)
	expected := []int{0, 1}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\"%v\" and \"%v\": expected result to be %v, got %v", nums, target, expected, actual)
	}
}
