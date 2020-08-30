/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	k := 0

	var v int
	for k, v = range m {
		if v == 1 {
			break
		}
	}

	return k
}

// @lc code=end

