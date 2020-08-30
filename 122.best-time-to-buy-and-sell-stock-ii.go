/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
const SEARCHING = 0
const HOLDING = 1

func maxProfit(prices []int) int {
	state := SEARCHING

	total := 0

	earn := 0
	buyPrice := 0
	for i := 0; i < len(prices)-1; i++ {
		if state == SEARCHING {
			if prices[i+1] > prices[i] {
				buyPrice = prices[i]
				state = HOLDING
			}
		} else {
			earn = max(earn, prices[i]-buyPrice)
			if prices[i+1] < prices[i] {
				state = SEARCHING
				total += earn
				earn = 0
			}
		}
	}

	if state == HOLDING {
		total += prices[len(prices)-1] - buyPrice
	}

	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

