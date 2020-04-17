package main

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count int
	parent, children := []*TreeNode{root}, []*TreeNode{}
	for len(parent) > 0 {
		count++
		children = []*TreeNode{}
		for _, v := range parent {
			if v.Left != nil {
				children = append(children, v.Left)
			}
			if v.Right != nil {
				children = append(children, v.Right)
			}
		}
		parent = children
	}
	return count
}
