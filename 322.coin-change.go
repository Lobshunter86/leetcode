/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp stores current best for target value
	dp := make([]int, amount+1)

	for target := range dp {
		for _, coin := range coins {
			// just match base case
			if target == coin {
				dp[target] = 1
				continue
			}

			left := target - coin
			if left >= 0 && dp[left] != 0 {
				if dp[target] != 0 {
					dp[target] = minInt(dp[target], dp[left]+1)
				} else {
					dp[target] = dp[left] + 1
				}
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}

	return dp[amount]
}

// should be fucking correct
// var memo = map[int]int{}

// func coinChange(coins []int, amount int) (result int) {
// 	if cached, ok := memo[amount]; ok {
// 		return cached
// 	}

// 	defer func() {
// 		memo[amount] = result
// 		println(amount, "returns: ", result)
// 	}()

// 	minCoins := -1
// 	for _, coin := range coins {
// 		if coin == amount {
// 			return 1
// 		} else if coin > amount {
// 			continue
// 		} else {
// 			minCoins = minPositive(minCoins, coinChange(coins, amount-coin))
// 		}
// 	}

// 	if minCoins > 0 {
// 		minCoins += 1
// 	}

// 	return minCoins
// }

// // minPositive won't have both arguments are lower than 0
// func minPositive(a, b int) int {
// 	if a <= 0 || a > b {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

