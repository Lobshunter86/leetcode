/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
var memo [][]int

func minDistance(word1 string, word2 string) int {
	if len(word1) < len(word2) {
		word1, word2 = word2, word1 // for cache friendly
	}

	memo = make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return lev(word1, word2, len(word1), len(word2))
}

// len(word1) >= len(word2)
func lev(word1, word2 string, i1, i2 int) (result int) {
	if min(i1, i2) == 0 {
		return max(i1, i2)
	}

	if v := memo[i1-1][i2-1]; v != -1 {
		return v
	}

	defer func() {
		memo[i1-1][i2-1] = result
	}()

	if word1[i1-1] == word2[i2-1] {
		return lev(word1, word2, i1-1, i2-1)
	}

	return 1 + min(
		lev(word1, word2, i1-1, i2), // delete

		min(lev(word1, word2, i1, i2-1), // insert
			lev(word1, word2, i1-1, i2-1)), // substitute
	)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

