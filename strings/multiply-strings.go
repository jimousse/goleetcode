package main

import (
	"strconv"
)

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			k, _ := strconv.Atoi(string(num1[i])) // 0-9
			p, _ := strconv.Atoi(string(num2[j])) // 0-9
			mul := k * p                          // 0-18
			res[i+j+1] += mul
			res[i+j] += res[i+j+1] / 10
			res[i+j+1] = res[i+j+1] % 10
		}
	}
	numZeros := 0
	for numZeros < len(res) && res[numZeros] == 0 {
		numZeros++
	}
	resStr := ""
	for i := numZeros; i < len(res); i++ {
		resStr += strconv.Itoa(res[i])
	}
	return resStr
}
