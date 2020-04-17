package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 *
 * https://leetcode.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (39.76%)
 * Total Accepted:    267.2K
 * Total Submissions: 670.5K
 * Testcase Example:  '1'
 *
 * The count-and-say sequence is the sequence of integers with the first five
 * terms as following:
 *
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 is read off as "one 1" or 11.
 * 11 is read off as "two 1s" or 21.
 * 21 is read off as "one 2, then one 1" or 1211.
 *
 * Given an integer n where 1 ≤ n ≤ 30, generate the n^th term of the
 * count-and-say sequence.
 *
 * Note: Each term of the sequence of integers will be represented as a
 * string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "1"
 *
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "1211"
 *
 */
func countAndSay2(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		array := []string{""}
		ch := string(s[0])
		start := 0
		count := 0
		for j := 0; j < len(s); j++ {
			if string(s[j]) == ch {
				count++
				array[start] = strconv.Itoa(count) + ch
			} else {
				count = 0
				start++
				ch = string(s[j])
				array = append(array, "")
				j--
			}
		}
		s = strings.Join(array, "")
	}
	return s
}
func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		var start int
		var ns string
		for j := 0; j < len(s); j++ {
			count := 1
			for start+count < len(s) && s[start+count] == s[start] {
				j = start + count
				count++
			}
			ns += strconv.Itoa(count) + string(s[start])
			start = j + 1
		}
		s = ns
	}
	return s
}
