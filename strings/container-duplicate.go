package main

// https://leetcode.com/problems/contains-duplicate/

// Given an integer array nums,
// return true if any value appears at least twice in the array,
// and return false if every element is distinct.

func ContainsDuplicate(num []int) bool {
	var elMap = make(map[int]bool)
	for _, el := range num {
		_, elFound := elMap[el]
		if elFound {
			return true
		} else {
			elMap[el] = true
		}
	}
	return false
}
