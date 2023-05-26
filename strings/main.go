package main

import "fmt"

func main() {

	// // unique email addresses
	// fmt.Println(
	// 	NumUniqueEmails([]string{
	// 		"test.email+alex@leetcode.com",
	// 		"test.e.mail+bob.cathy@leetcode.com",
	// 		"test.e.mail+bob.cathy@leetcode.com",
	// 		"testemail+david@lee.tcode.com",
	// 	}))

	// // contains duplicate
	// fmt.Println(ContainsDuplicate([]int{
	// 	1, 2, 3, 4, 6, 8, 4, 6, 4, 1, 3, 4,
	// }))

	// // is valid anagram
	// fmt.Println(IsAnagram("rat", "car"))

	// // check s2 contains a permutation of s1
	// fmt.Println(CheckInclusion("adc", "dcda"))

	// fmt.Println(CanPermutePalindrome("carerac"))

	// fmt.Println(IsOneEditDistance("a", "ac"))

	// fmt.Println(MinDistance("horse", "ros"))

	// matrixToRotate := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// RotateImage(matrixToRotate)
	// fmt.Println(matrixToRotate)

	// fmt.Println(RotateString("abcde", "cdeab"))

	// matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	// matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix := [][]int{{0}, {1}}
	SetMatrixZeroes(matrix)
	fmt.Println(matrix)
}
