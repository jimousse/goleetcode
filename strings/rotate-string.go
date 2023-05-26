package main

import "strings"

// https://leetcode.com/problems/rotate-string/

func RotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
