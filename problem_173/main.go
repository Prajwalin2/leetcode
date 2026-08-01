/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	arr     []int
	current int
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		arr:     []int{},
		current: 0,
	}
	if root != nil {
		iter.arr = append(iter.arr, Constructor(root.Left).arr...)
		iter.arr = append(iter.arr, root.Val)
		iter.arr = append(iter.arr, Constructor(root.Right).arr...)
	}
	return iter
}

func (this *BSTIterator) Next() int {
	this.current++
	return this.arr[this.current]
}

func (this *BSTIterator) HasNext() bool {
	return this.current+1 < len(this.arr)
}
