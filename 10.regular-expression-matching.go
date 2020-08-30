/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
var memo [][]int

const UNKNOWN = 0
const TRUE = 1
const FALSE = 2

func isMatch(s string, p string) bool {
	memo = make([][]int, len(s)+2)
	for i := range memo {
		memo[i] = make([]int, len(p)+2)
	}

	return dp(s, p, 0, 0) == TRUE
}

func dp(s, p string, si, pi int) (result int) {
	if memo[si][pi] != UNKNOWN {
		return memo[si][pi]
	}

	defer func() {
		memo[si][pi] = result
	}()

	if si == len(s) && pi == len(p) {
		return TRUE
	}

	if pi == len(p) {
		return FALSE
	}

	if si == len(s) {
		if (len(p)-pi)%2 != 0 {
			return FALSE
		}

		for pi++; pi < len(p); pi += 2 {
			if p[pi] != '*' {
				return FALSE
			}
		}
		return TRUE
	}

	if s[si] == p[pi] || p[pi] == '.' {
		if pi < len(p)-1 && p[pi+1] == '*' {
			if (dp(s, p, si+1, pi) == TRUE) ||
				(dp(s, p, si, pi+2) == TRUE) {
				return TRUE
			}
			return FALSE
		} else {
			return dp(s, p, si+1, pi+1)
		}
	} else {
		if pi < len(p)-1 && p[pi+1] == '*' {
			return dp(s, p, si, pi+2)
		} else {
			return FALSE
		}
	}
}

// @lc code=end

