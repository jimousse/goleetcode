package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	input := "dvdf"
	actual := LengthOfLongestSubstring(input)
	expected := 3

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = "bbbbb"
	actual = LengthOfLongestSubstring(input)
	expected = 1

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = "pwwkew"
	actual = LengthOfLongestSubstring(input)
	expected = 3

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = " "
	actual = LengthOfLongestSubstring(input)
	expected = 1

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}

	input = "abba"
	actual = LengthOfLongestSubstring(input)
	expected = 2

	if actual != expected {
		t.Errorf("\"%v\": expected result to be %v, got %v", input, expected, actual)
	}
}
