package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Implement by hash map
	hashMap := map[int]int{}
	for i, v := range nums {
		hashMap[v] = i
	}
	for i, v := range nums {
		rest := target - v
		if j, ok := hashMap[rest]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
}
