/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
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

// leetcode version
type Codec struct {
	builder strings.Builder
	buffer  []string
}

func Constructor() Codec {
	return Codec{
		builder: strings.Builder{},
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// this.levelOrder(root)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	var node *TreeNode
	for len(queue) != 0 {
		node, queue = queue[0], queue[1:]

		this.builder.Write([]byte(","))
		if node == nil {
			this.builder.Write([]byte("null"))
			continue
		}

		this.builder.Write([]byte(strconv.Itoa(node.Val)))
		queue = append(queue, node.Left, node.Right)
	}

	result := this.builder.String()[1:]
	for strings.HasSuffix(result, ",null") {
		result = result[:len(result)-len(",null")]
	}

	return "[" + result + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[1 : len(data)-1]

	this.buffer = strings.Split(data, ",")
	if this.buffer[0] == "null" {
		return nil
	}

	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(this.buffer[0])
	this.buffer = this.buffer[1:]

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	var node *TreeNode
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]
		var left, right *TreeNode

		// left
		if len(this.buffer) == 0 {
			break
		}
		if this.buffer[0] != "null" {
			left = &TreeNode{}
			left.Val, _ = strconv.Atoi(this.buffer[0])
			queue = append(queue, left)
		}
		this.buffer = this.buffer[1:]
		node.Left = left

		// right
		if len(this.buffer) == 0 {
			break
		}
		if this.buffer[0] != "null" {
			right = &TreeNode{}
			right.Val, _ = strconv.Atoi(this.buffer[0])
			queue = append(queue, right)
		}
		this.buffer = this.buffer[1:]
		node.Right = right
	}

	return root
}

// type Codec struct {
// 	builder strings.Builder
// 	buffer  []string
// }

// func Constructor() Codec {
// 	return Codec{
// 		builder: strings.Builder{},
// 	}
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
// 	this.postorder(root)
// 	return this.builder.String()[1:]
// }

// func (this *Codec) postorder(root *TreeNode) {
// 	if root == nil {
// 		this.builder.Write([]byte(",null"))
// 		return
// 	}

// 	this.postorder(root.Left)
// 	this.postorder(root.Right)
// 	this.builder.Write([]byte(","))
// 	this.builder.Write([]byte(strconv.Itoa(root.Val)))
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	this.buffer = strings.Split(data, ",")
// 	return this.restore()
// }

// func (this *Codec) restore() *TreeNode {
// 	var val string
// 	val, this.buffer = this.buffer[len(this.buffer)-1], this.buffer[:len(this.buffer)-1]

// 	if val == "null" {
// 		return nil
// 	}

// 	node := &TreeNode{}
// 	node.Val, _ = strconv.Atoi(val)
// 	node.Right = this.restore()
// 	node.Left = this.restore()

// 	return node
// }

// preorder tranverse
// type Codec struct {
// 	builder strings.Builder
// 	buffer  []string
// }

// func Constructor() Codec {
// 	return Codec{
// 		builder: strings.Builder{},
// 	}
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
// 	this.preorder(root)
// 	return this.builder.String()[1:]
// }

// func (this *Codec) preorder(root *TreeNode) {
// 	if root == nil {
// 		this.builder.Write([]byte(",null"))
// 		return
// 	}

// 	this.builder.Write([]byte(","))
// 	this.builder.Write([]byte(strconv.Itoa(root.Val)))
// 	this.preorder(root.Left)
// 	this.preorder(root.Right)
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	this.buffer = strings.Split(data, ",")
// 	return this.restore()
// }

// func (this *Codec) restore() *TreeNode {
// 	var val string
// 	val, this.buffer = this.buffer[0], this.buffer[1:]

// 	if val == "null" {
// 		return nil
// 	}

// 	node := &TreeNode{}
// 	node.Val, _ = strconv.Atoi(val)
// 	node.Left = this.restore()
// 	node.Right = this.restore()
// 	return node
// }

//  VERSION R U KINDING ME
// type Codec struct {
// 	cache map[string]*TreeNode
// }

// func Constructor() Codec {
// 	return Codec{cache: make(map[string]*TreeNode)}
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
// 	id := time.Now().String()
// 	this.cache[id] = root
// 	return id
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	return this.cache[data]
// }

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

