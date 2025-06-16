package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = nums2
	}
	ptr1 := m - 1
	ptr2 := n - 1
	ptr3 := m + n - 1
	for ptr3 >= 0 && ptr1 >= 0 && ptr2 >= 0 {
		if nums1[ptr1] > nums2[ptr2] {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		} else {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		}
		ptr3--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
