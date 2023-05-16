package main

// https://leetcode.com/problems/valid-anagram/description/

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count = make(map[rune]int)

	for _, v := range s {
		count[v]++
	}

	for _, v := range t {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}

	return true
}
