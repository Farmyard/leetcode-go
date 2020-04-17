package main

import "fmt"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (42.51%)
 * Total Accepted:    1.6M
 * Total Submissions: 3.7M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 *
 *
 */
func twoSum(nums []int, target int) []int {
	var ret []int
	l := len(nums)
	for i := 0; i < l; i++ {
		// if nums[i] > target {
		// 	continue
		// }
		for j := i + 1; j < l; j++ {
			// if nums[j] > target {
			// 	continue
			// }
			fmt.Println(i, j)
			if nums[i]+nums[j] == target {
				ret = []int{i, j}
			}
		}
	}
	return ret
}

func twoSum2(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		if idx, ok := mp[target-num]; ok {
			return []int{idx, i}
		}
		mp[num] = i
	}
	return nil
}
