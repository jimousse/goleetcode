package main

func MaxWaterArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left, right := 0, len(height)-1
	maxArea := (right - left) * min(height[right], height[left])

	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		maxArea = max((right-left)*min(height[right], height[left]), maxArea)
	}
	return maxArea
}
