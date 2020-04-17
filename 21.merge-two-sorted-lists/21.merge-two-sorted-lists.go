package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (46.17%)
 * Total Accepted:    530.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	b := l
	for ; l1 != nil; l1 = l1.Next {
		for ; l2 != nil && l1.Val > l2.Val; l2 = l2.Next {
			l.Next = l2
			l = l.Next
		}
		l.Next = l1
		l = l.Next
	}
	l.Next = l2
	b = b.Next
	return b
}
