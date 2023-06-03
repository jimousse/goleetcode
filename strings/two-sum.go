package main

// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	// map of complement (target - n) to index in array
	compMap := make(map[int]int)

	for i, n := range nums {
		comp := target - n
		_, found := compMap[comp]
		if found {
			return []int{compMap[comp], i}
		}
		compMap[n] = i
	}
	return []int{}
}
