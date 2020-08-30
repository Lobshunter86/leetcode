/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	memo := make([]int, 0, len(nums))
	memo = append(memo, 1)
	for i := 1; i < len(nums); i++ {
		maxLen := 1
		for j := range memo {
			if nums[i] > nums[j] {
				maxLen = max(maxLen, memo[j]+1)
			}
		}

		memo = append(memo, maxLen)
	}

	m := memo[0]
	for _, v := range memo {
		m = max(m, v)
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end