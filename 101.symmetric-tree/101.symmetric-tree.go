package main

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := []*TreeNode{root.Left, root.Right}
	i, j := 0, len(l)-1
	for j > i {
		n := i
		for j > i {
			if l[i] == l[j] {
				i++
				j--
				continue
			}
			if (l[i] == nil && l[j] != nil) || (l[i] != nil && l[j] == nil) || (l[i].Val != l[j].Val) {
				return false
			}
			i++
			j--
		}

		i = len(l)
		for n < i {
			if l[n] != nil {
				l = append(l, l[n].Left, l[n].Right)
			}
			n++
		}
		j = len(l) - 1
	}

	return true
}
