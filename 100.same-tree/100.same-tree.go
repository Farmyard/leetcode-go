package main

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree1(p *TreeNode, q *TreeNode) bool {
	P := []*TreeNode{p}
	Q := []*TreeNode{q}
	for len(P) > 0 && len(Q) > 0 {
		tmpP, P := P[0], P[1:]
		tmpQ, Q := Q[0], Q[1:]
		if tmpP == nil {
			if tmpQ != nil {
				return false
			}
			continue
		}
		if tmpQ == nil {
			if tmpP != nil {
				return false
			}
			continue
		}
		if tmpP.Val != tmpQ.Val {
			return false
		} else {
			P = append(P, tmpP.Left, tmpP.Right)
			Q = append(Q, tmpQ.Left, tmpQ.Right)
		}
	}
	return len(P) == len(Q)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	P := []*TreeNode{p}
	Q := []*TreeNode{q}
	for len(P) > 0 && len(Q) > 0 {
		p, P = P[0], P[1:]
		q, Q = Q[0], Q[1:]
		if p == nil {
			if q != nil {
				return false
			}
			continue
		}
		if q == nil {
			if p != nil {
				return false
			}
			continue
		}
		if p.Val != q.Val {
			return false
		} else {
			P = append(P, p.Left, p.Right)
			Q = append(Q, q.Left, q.Right)
		}
	}
	return len(P) == len(Q)
}
