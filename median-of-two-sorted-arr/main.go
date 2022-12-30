package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arrJoin := append(nums1, nums2...)
	sort.Ints(arrJoin)
	lenArr := len(arrJoin)
	if lenArr % 2 == 1 {
		return float64(arrJoin[(lenArr - 1) / 2])
	} else {
		return float64(arrJoin[lenArr / 2] + arrJoin[lenArr / 2 - 1]) / 2
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
