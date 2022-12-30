package numbersteptozero

import (
	"fmt"
	"strings"
)

// func numberOfSteps(num int) int {
// 	if num == 0 {
// 		return 0
// 	}
// 	if num%2 == 0 {
// 		return numberOfSteps(num/2) + 1
// 	} else {
// 		return 1 + numberOfSteps(num-1)
// 	}
// }

func numberOfSteps(num int) int {
	digits := fmt.Sprintf("%b", num)
	return len(digits) + strings.Count(digits, "1") - 1
}
