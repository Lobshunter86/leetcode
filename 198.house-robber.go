/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	memo := make([]int, len(nums))
	memo[0] = nums[0]
	memo[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		for j := 0; j <= i-2; j++ {
			memo[i] = max(memo[i], nums[i]+memo[j])
		}
	}

	result := 0
	for _, v := range memo {
		result = max(result, v)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

