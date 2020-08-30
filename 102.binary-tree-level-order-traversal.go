/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curLevel := make([]int, 0, len(queue))
		newQueue := make([]*TreeNode, 0, len(queue)*2)

		for _, node := range queue {
			// if node == nil {
			// 	continue
			// }
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}

			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}

		result = append(result, curLevel)
		queue = newQueue
	}

	return result
}

// @lc code=end

