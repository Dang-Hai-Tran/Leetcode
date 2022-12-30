package main
import (
	"fmt"
	"strings"
)

func lenOfLongestSubstring(s string) int {
	countMax := 1
	for i := 0; i < len(s) - 1; i++ {
		count := 1
		for j := i + 1; j < len(s); j++ {
			if !strings.Contains(s[i:j], string(s[j])) {
				count++
			} else {
				break
			}
		}
		if (count > countMax) {
			countMax = count
		}
	}
	return countMax
}

func main() {
	s := "pwwkew"
	fmt.Println("longest :", lenOfLongestSubstring(s))
}