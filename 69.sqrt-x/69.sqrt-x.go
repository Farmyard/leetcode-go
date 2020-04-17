package main

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (30.86%)
 * Total Accepted:    345.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */
func mySqrt(x int) int {
	start, end := 1, x/2
	if end == 0 {
		return x
	}
	start, end = find(end/2, end, x)
	for {
		mid := (start + end) / 2
		if mid*mid > x {
			end = mid
		} else if mid*mid < x {
			start = mid
		} else {
			return mid
		}
		if start == end {
			return start
		} else if start+1 == end {
			if end*end <= x {
				return end
			}
			return start
		}
	}
}

func find(s int, e int, x int) (int, int) {
	if (s+e)/2*(s+e)/2 > x {
		s, e = find(s/2, (s+e)/2, x)
	}
	return s, e
}
