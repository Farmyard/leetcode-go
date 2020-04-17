package main

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (31.43%)
 * Total Accepted:    394.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */
func strStr2(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	var n []string
	var J int
	for i := 0; i < len(needle); i++ {
		tmp := needle[i]

		for j := 0; i < len(needle) && j < len(haystack); {

			if tmp == haystack[j] {
				if len(n) == 0 {
					J = j
				}
				n = append(n, string(tmp))
				if i+1 >= len(needle) {
					return J
				}
				tmp = needle[i+1]
				i++
				j++
				if j > len(haystack)-1 {
					return -1
				}
			} else if len(n) > 0 && tmp != haystack[j] {
				n = []string{}
				i = 0
				tmp = needle[i]
				j = J + 1
			} else if j+len(needle) >= len(haystack) {
				return -1
			} else {
				j++
			}
		}
	}
	return J
}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	J := -1
	for i, j := 0, 0; i < len(needle) && j < len(haystack); j++ {
		if needle[i] == haystack[j] {
			if J == -1 {
				J = j
			}
			i++
			if i < len(needle) && j == len(haystack)-1 {
				return -1
			}
		} else {
			i = 0
			if J >= 0 {
				j = J
			}
			J = -1
		}
	}
	return J
}
