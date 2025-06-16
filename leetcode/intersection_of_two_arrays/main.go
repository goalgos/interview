package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]struct{}, len(nums1))
	var result []int
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	for _, num := range nums1 {
		map1[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := map1[num]; ok {
			result = append(result, num)
			delete(map1, num)
		}
	}
	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
