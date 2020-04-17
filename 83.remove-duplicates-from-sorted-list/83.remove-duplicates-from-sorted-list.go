package main

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (42.05%)
 * Total Accepted:    312.3K
 * Total Submissions: 740.8K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates1(head *ListNode) *ListNode {
	l := &ListNode{}
	b := l
	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		l.Next = head
		l = l.Next
		head = head.Next
	}

	return b.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	l := &ListNode{}
	b := l
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
			l.Next = head
		} else {
			if l.Next == nil {
				l.Next = head
			}
			l = l.Next
			head = head.Next
		}
	}
	return b.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	l := &ListNode{}
	b := l
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
			l.Next = head
		} else {
			l.Next = head
			l = l.Next
			head = head.Next
		}
	}
	return b.Next
}
