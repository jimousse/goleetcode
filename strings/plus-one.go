package main

// https://leetcode.com/problems/plus-one/description/

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			if i > 0 {
				break
			}
		}
	}
	return digits
}
