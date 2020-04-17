package main

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var array [][]int
	if root == nil {
		return array
	}
	parent, children := []*TreeNode{root}, []*TreeNode{}
	for len(parent) > 0 {
		children = []*TreeNode{}
		arr := []int{}
		for _, v := range parent {
			arr = append(arr, v.Val)
			if v.Left != nil {
				children = append(children, v.Left)
			}
			if v.Right != nil {
				children = append(children, v.Right)
			}
		}
		array = append(array, arr)
		parent = children
	}
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-i-1] = array[len(array)-i-1], array[i]
	}
	return array
}
