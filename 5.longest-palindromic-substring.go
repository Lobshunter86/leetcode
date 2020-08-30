/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var result string
	for i := range s {
		result = longestStr(
			result,
			longestStr(
				palidrome(s, i, i+1),
				palidrome(s, i, i),
			),
		)
	}

	return result
}

func palidrome(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}

	return s[left+1 : right]
}

func longestStr(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

// @lc code=end

