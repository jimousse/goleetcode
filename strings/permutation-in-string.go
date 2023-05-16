package main

// Given two strings s1 and s2,
// return true if s2 contains a permutation of s1,
// or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.

// https://leetcode.com/problems/permutation-in-string/description/

func matches(s1Map [26]int, s2Map [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1Map[i] != s2Map[i] {
			return false
		}
	}
	return true
}

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map, s2Map := [26]int{}, [26]int{}

	// create the hashmap for the first window
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	// update the window, when sliding
	for i := 0; i < len(s2)-len(s1); i++ {
		if matches(s1Map, s2Map) {
			return true
		}
		s2Map[s2[i]-'a']--
		s2Map[s2[i+len(s1)]-'a']++
	}

	return matches(s1Map, s2Map)
}
