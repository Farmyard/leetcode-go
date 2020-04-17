package main

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
func merge1(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[n+m-i-1] = nums2[n-i-1]
	}
	i := 0
	for i < n {
		tmp := nums1[m+i]
		j := m + i - 1
		for j >= 0 {
			if nums1[j] > tmp {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			} else {
				break
			}
			j--
		}
		i++
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	for i := 0; i < n; i++ {
		nums1[n+m-i-1] = nums2[n-i-1]
	}
	i := m
	j := i - 1
	tmp := nums1[i]
	for i < n+m {
		if j >= 0 && nums1[j] > tmp {
			nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			j--
		} else {
			i++
			j = i - 1
			if i < n+m {
				tmp = nums1[i]
			}
		}
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	i, j := 0, m
	for i < n {
		if j-1 >= 0 && nums2[i] < nums1[j-1] {
			nums1[j] = nums1[j-1]
			j--
		} else {
			nums1[j] = nums2[i]
			i++
			j = m + i
		}
	}
}
