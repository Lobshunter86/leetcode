/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	earn := 0
	buy := prices[0]

	for _, v := range prices {
		buy, earn = min(buy, v), max(earn, v-buy)
	}

	return earn
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

