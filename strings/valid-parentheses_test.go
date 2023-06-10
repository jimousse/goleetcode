package main

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	input := "()"
	actual := IsValidParentheses(input)
	expected := true

	if actual != expected {
		t.Errorf("\"%v\": expected result to be \"%v\", got \"%v\"", input, expected, actual)
	}

	input = "()[]{}"
	actual = IsValidParentheses(input)
	expected = true

	if actual != expected {
		t.Errorf("\"%v\": expected result to be \"%v\", got \"%v\"", input, expected, actual)
	}

	input = "(]"
	actual = IsValidParentheses(input)
	expected = false

	if actual != expected {
		t.Errorf("\"%v\": expected result to be \"%v\", got \"%v\"", input, expected, actual)
	}

	input = "]]"
	actual = IsValidParentheses(input)
	expected = false

	if actual != expected {
		t.Errorf("\"%v\": expected result to be \"%v\", got \"%v\"", input, expected, actual)
	}

}
