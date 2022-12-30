package sum1darray

// "fmt"

// func sumArray(nums ...int) int {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	return sum
// }

// func runningSum(nums []int) []int {
// 	res := []int{}
// 	for i := range nums {
// 		res = append(res, sumArray(nums[:i+1]...))
// 	}
// 	return res
// }

// func runningSum(nums []int) []int {
// 	res := make([]int, len(nums))
// 	for i := range res {
// 		if i == 0 {
// 			res[i] = nums[i]
// 		} else {
// 			res[i] = res[i-1] + nums[i]
// 		}
// 	}
// 	return res
// }

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
