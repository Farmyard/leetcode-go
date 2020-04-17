package main

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.11%)
 * Total Accepted:    424.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return ""
		}
		return strs[0]
	}
	i, j := 1, 0
	var s string
	for j < len(strs[i-1]) && j < len(strs[i]) && string(strs[i-1][j]) == string(strs[i][j]) {
		i++
		if i == len(strs) {
			s += string(strs[i-1][j])
			i = 1
			j++
		}
	}
	return s
}
