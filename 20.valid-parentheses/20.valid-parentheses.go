package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (36.04%)
 * Total Accepted:    536.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */
func isValid(s string) bool {
	mp := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	var left []string
	for _, v := range s {
		if _, ok := mp[string(v)]; ok {
			if len(left) > 0 && left[len(left)-1] == mp[string(v)] {
				left = left[:len(left)-1]
			} else {
				return false
			}
		} else {
			left = append(left, string(v))
		}
	}

	return len(left) == 0
}
