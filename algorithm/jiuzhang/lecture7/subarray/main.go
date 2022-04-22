package main

import "math"

// 子数组
// 最佳股票销售时间
// 子数组 1，2，3，4

// 最小子数组
// 给定一个含有n个正整数的数组和一个正整数 target 。
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。


// 滑动窗口
//	处理连续子数组

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ret := math.MaxInt32
	start, end := 0, 0
	sum := 0

	for end < len(nums) {
		sum = sum + nums[end]
		for sum >= target {
			ret = min(ret, end - start + 1) // 满足条件的最小子数组
			sum = sum - nums[start]
			start++
		}
		end++
	}

	if ret == math.MaxInt32 {
		return 0
	}

	return ret
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}