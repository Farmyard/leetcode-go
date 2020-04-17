package main

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (40.55%)
 * Total Accepted:    375.4K
 * Total Submissions: 925.2K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */
func searchInsert2(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return i
}
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if target <= nums[left] {
		return left
	}
	for right-left >= 0 {
		mid := (right + left) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
