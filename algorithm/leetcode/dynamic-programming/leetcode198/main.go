package main


import "fmt"

func rob(nums []int) int {
	fmt.Println(nums)
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	dp := make([]int, l)


	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < l; i++ {
		dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])
	}

	fmt.Println(dp)
	return dp[l - 1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	ns := []int{1, 2, 3, 1}
	rob(ns)
}