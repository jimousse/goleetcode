package main

import "sort"

// https://leetcode.com/problems/3sum/description/
// returns all triplets [nums[i], nums[j], nums[k]]
// such as nums[i] + nums[j] + nums[k] == 0

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		// look for two-sum -n in the rest
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
