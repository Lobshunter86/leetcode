/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var memo map[*TreeNode]int

func rob(root *TreeNode) int {
	memo = make(map[*TreeNode]int)

	dp(root)

	m := 0
	for _, v := range memo {
		m = max(m, v)
	}

	return m
}

func dp(root *TreeNode) (result int) {
	if r, ok := memo[root]; ok {
		return r
	}

	defer func() {
		memo[root] = result
	}()

	if root == nil {
		return 0
	}

	leftGrand, rightGrand := 0, 0
	left, right := dp(root.Left), dp(root.Right)
	if root.Left != nil {
		leftGrand = dp(root.Left.Left) + dp(root.Left.Right)
	}
	if root.Right != nil {
		rightGrand = dp(root.Right.Left) + dp(root.Right.Right)
	}

	return max(root.Val+leftGrand+rightGrand, left+right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

