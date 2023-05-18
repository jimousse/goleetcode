package main

// https://leetcode.com/problems/palindrome-permutation/description/

func CanPermutePalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	// create hashtable of letters
	var letterMap = [26]int{}
	for _, v := range s {
		letterMap[v-'a']++
	}

	var numOddFrequency int

	for _, v := range letterMap {
		if v%2 == 1 {
			numOddFrequency++
		}
	}

	return numOddFrequency <= 1
}
