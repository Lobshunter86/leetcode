/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	memo := make([]int, len(nums))
	memo[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if memo[i-1] > 0 {
			memo[i] = nums[i] + memo[i-1]
		} else {
			memo[i] = nums[i]
		}
	}

	maxSum := memo[0]
	for _, v := range memo[1:] {
		maxSum = max(maxSum, v)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

