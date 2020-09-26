/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make([]bool, 256)

	longest := 0
	start := 0
	for i, r := range s {
		if m[r] {
			longest = max(i-start, longest)
			for j := start; ; j++ {
				if s[j] == byte(r) {
					start = j + 1
					break
				}

				m[s[j]] = false
			}

		} else {
			m[r] = true
		}
	}

	longest = max(len(s)-start, longest)

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

