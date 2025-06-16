package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	tmpCount := 0
	maxCount := 0
	for _, num := range nums {
		if num == 1 {
			tmpCount++
			maxCount = max(maxCount, tmpCount)
		} else {
			tmpCount = 0
		}
	}

	return maxCount
}

func main() {
	result := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	fmt.Println(result)
}
