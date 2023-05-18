package main

import (
	"math"
)

// https://leetcode.com/problems/one-edit-distance/description/

func IsOneEditDistance(s string, t string) bool {
	if math.Abs(float64(len(s)-len(t))) > 1 || s == t {
		return false
	}

	var diffFound bool
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			if diffFound {
				return false
			}
			diffFound = true
			if len(s) == len(t) {
				i++
				j++
			} else if len(s) < len(t) {
				j++
			} else {
				i++
			}
		}
	}

	return diffFound || i < len(s) || j < len(t)
}
