/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	maxVal := nums[0]
	maxCur, minCur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			maxCur, minCur = max(maxCur*nums[i], nums[i]), min(minCur*nums[i], nums[i])
		} else {
			// cautious here, split into 2 lines will lead to wrong result
			// because maxCur is updated, but minCur = ... needs the old one
			maxCur, minCur = max(minCur*nums[i], nums[i]), min(maxCur*nums[i], nums[i])
		}

		maxVal = max(maxVal, maxCur)
	}

	return maxVal
}

// version1 intuitive
// func maxProduct(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	// maxP, minP := nums[0], nums[0]
// 	// cur := nums[0]
// 	maxP, minP := make([]int, len(nums)), make([]int, len(nums))
// 	maxP[0], minP[0] = nums[0], nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] > 0 {
// 			maxP[i] = max(maxP[i-1]*nums[i], nums[i])
// 			minP[i] = min(minP[i-1]*nums[i], nums[i])
// 		} else {
// 			maxP[i] = max(minP[i-1]*nums[i], nums[i])
// 			minP[i] = min(maxP[i-1]*nums[i], nums[i])
// 		}
// 	}

// 	m := maxP[0]
// 	for _, v := range maxP {
// 		m = max(m, v)
// 	}
// 	return m
// }

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

