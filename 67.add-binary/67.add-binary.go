package main

import (
	"strconv"
)

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (38.22%)
 * Total Accepted:    286.9K
 * Total Submissions: 746.9K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	i := len(b) - 1
	j := len(a) - 1
	var c, s int
	var ret string
	for j >= 0 {
		bb := 0
		aa, _ := strconv.Atoi(string(a[j]))
		if i >= 0 {
			bb, _ = strconv.Atoi(string(b[i]))
		}
		c, s = halfAdder(aa, bb, c)
		ret = strconv.Itoa(s) + ret
		i--
		j--
	}
	if c > 0 {
		ret = "1" + ret
	}
	return ret
}

func halfAdder(a int, b int, c int) (C int, S int) {
	if ((a ^ b) ^ c) == 1 {
		S = 1
	} else {
		S = 0
	}
	if ((a & b) ^ ((a ^ b) & c)) == 1 {
		C = 1
	} else {
		C = 0
	}
	return C, S
}
