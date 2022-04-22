package main

import "fmt"

// 考虑环形的，分为 [0, len - 2], [1, len - 1] 2种情况，没别求最大偷取

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	one := _rob(nums[:len(nums) - 1])
	two := _rob(nums[1:])

	return max(one, two)
}


func _rob(nums []int) int {
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
	ns := []int{2, 3, 2}
	fmt.Println(rob(ns))
}