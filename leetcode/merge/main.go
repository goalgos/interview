package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	for _, cur := range intervals[1:] {
		previous := result[len(result)-1]
		if cur[0] < previous[1] {
			previous[1] = max(previous[1], cur[1])
		} else {
			result = append(result, cur)
		}
	}

	return result
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
