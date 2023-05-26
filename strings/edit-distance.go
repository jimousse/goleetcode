package main

func MinDistance(word1 string, word2 string) int {
	mapSize := max(len(word1), len(word2)) + 1
	var memo = make(map[[2]int]int, mapSize)
	return minDistanceRecur(word1, word2, len(word1), len(word2), memo)
}

func minDistanceRecur(word1 string, word2 string, word1Index int, word2Index int, memo map[[2]int]int) int {
	if word1Index == 0 {
		return word2Index
	}

	if word2Index == 0 {
		return word1Index
	}

	pair := [2]int{word1Index, word2Index}
	v, found := memo[pair]
	if found {
		return v
	}

	if word1[word1Index-1] == word2[word2Index-1] {
		minDistance := minDistanceRecur(word1, word2, word1Index-1, word2Index-1, memo)
		memo[pair] = minDistance
		return minDistance
	} else {
		replaceOp := minDistanceRecur(word1, word2, word1Index-1, word2Index-1, memo)
		insertOp := minDistanceRecur(word1, word2, word1Index, word2Index-1, memo)
		deleteOp := minDistanceRecur(word1, word2, word1Index-1, word2Index, memo)
		minDistance := min(replaceOp, min(insertOp, deleteOp)) + 1
		memo[pair] = minDistance
		return minDistance
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
