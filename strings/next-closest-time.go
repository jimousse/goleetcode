package main

import (
	"strconv"
)

// https://leetcode.com/problems/next-closest-time/description/

func NextClosestTime(time string) string {
	numMap := make(map[int]bool, 4)
	currHour, _ := strconv.Atoi(time[0:2])
	currMin, _ := strconv.Atoi(time[3:])
	numMap[currHour/10] = true
	numMap[currHour%10] = true
	numMap[currMin/10] = true
	numMap[currMin%10] = true
	var result string
	for {
		currMin++
		if currMin == 60 {
			currHour++
			currMin = 0
		}
		if currHour == 24 {
			currHour = 0
		}

		digit1 := currHour / 10
		digit2 := currHour % 10
		digit3 := currMin / 10
		digit4 := currMin % 10
		allContained := numMap[digit1] && numMap[digit2] && numMap[digit3] && numMap[digit4]
		if allContained {
			result = strconv.Itoa(digit1) + strconv.Itoa(digit2) + ":" + strconv.Itoa(digit3) + strconv.Itoa(digit4)
			break
		}
	}
	return result
}
