package main

import (
	"math"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.21%)
 * Total Accepted:    633.2K
 * Total Submissions: 2.5M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */
func reverse(x int) int {
	var arr []int
	for X := int(math.Abs(float64(x))); X > 0; {
		X = X / 10
		arr = append(arr, X%10)
	}

	var sum int
	for i := 0; i < len(arr); i++ {
		if x < 0 {
			sum -= arr[i] * int(math.Pow10(len(arr)-i-1))
		} else {
			sum += arr[i] * int(math.Pow10(len(arr)-i-1))
		}
	}

	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}
	return sum
}

func reverse2(x int) int {
	ans := 0
	for x != 0 {
		last := x % 10
		x /= 10
		next := ans*10 + last
		if (next-last)/10 != ans || next > math.MaxInt32 || next < math.MinInt32 {
			return 0
		}
		ans = next
	}

	return ans
}

func reverse3(x int) int {
	s := 0
	for x != 0 {
		y := x % 10
		x /= 10
		s = s*10 + y
		if s > math.MaxInt32 || s < math.MinInt32 {
			return 0
		}
	}
	return s
}
