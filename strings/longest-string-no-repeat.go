package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func LengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[byte]int)
	ans := 0
	for i, j := 0, 0; j < len(s); j++ {
		j_bis, found := charIndexMap[s[j]]
		if found {
			if i < j_bis+1 {
				i = j_bis + 1
			}
		}
		if ans < j-i+1 {
			ans = j - i + 1
		}
		charIndexMap[s[j]] = j
	}
	return ans
}
