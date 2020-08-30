/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
var memo []int

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}

	memo = make([]int, n+1)
	memo[0] = 1
	memo[1] = 1

	return dp(n)
}

func dp(n int) (result int) {
	if memo[n] != 0 {
		return memo[n]
	}

	defer func() {
		memo[n] = result
	}()

	result = 0
	for root := 1; root <= n; root++ {
		result += dp(root-1) * dp(n-root)
	}

	return result
}

// @lc code=end

