package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	input1 := "123"
	input2 := "456"
	actual := Multiply(input1, input2)
	expected := "56088"

	if actual != expected {
		t.Errorf("\"%v\" * \"%v\": expected result to be \"%v\", got \"%v\"", input1, input2, expected, actual)
	}

	input1 = "498828660196"
	input2 = "840477629533"
	actual = Multiply(input1, input2)
	expected = "419254329864656431168468"

	if actual != expected {
		t.Errorf("\"%v\" * \"%v\": expected result to be \"%v\", got \"%v\"", input1, input2, expected, actual)
	}
}
