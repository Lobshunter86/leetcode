/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func isValidBST(root *TreeNode) bool {
	return isValid(root, MinInt, MaxInt)
}

func isValid(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}

	if root.Left != nil &&
		root.Left.Val >= root.Val {
		return false
	}

	if root.Right != nil &&
		root.Right.Val <= root.Val {
		return false
	}

	return isValid(root.Left, minVal, root.Val) && isValid(root.Right, root.Val, maxVal)
}

// @lc code=end

