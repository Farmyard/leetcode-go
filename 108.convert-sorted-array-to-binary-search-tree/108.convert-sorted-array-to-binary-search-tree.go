package main

import "fmt"

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	for _, v := range nums {
		makeTree(v, root)
	}
	fmt.Println(root)
	return root
}

func makeTree(v int, root *TreeNode) {
	if root == nil {
		root = &TreeNode{
			Val: v,
		}
	}
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(nums)
}
