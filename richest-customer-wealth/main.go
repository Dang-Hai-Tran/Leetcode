package richestcustomerwealth

func sumSlice(sl []int) int {
	sum := 0
	for i := range sl {
		sum += sl[i]
	}
	return sum
}
func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for i := range accounts {
		wealth := sumSlice(accounts[i])
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
