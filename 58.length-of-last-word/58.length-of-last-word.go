package main

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.18%)
 * Total Accepted:    254.6K
 * Total Submissions: 790.8K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 */
func lengthOfLastWord(s string) int {
	current, count := 0, 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " && count > 0 {
			current = count
			count = 0
		}
		if string(s[i]) != " " {
			count++
		}
	}
	if count > 0 {
		return count
	}
	return current
}
