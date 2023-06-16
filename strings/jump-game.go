package main

// https://leetcode.com/problems/jump-game/description/

func CanJump(nums []int) bool {
	lastGoodPosition := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= lastGoodPosition {
			lastGoodPosition = i
		}
	}
	return lastGoodPosition == 0
}
