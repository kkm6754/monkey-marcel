package main

import "fmt"

// 二分法
// 给一个已经排序的整数数组，找 任何位置、第一个位置、最后一个位置 的target
// 没有返回 -1 结束

// 递归 Recursion (消耗系统栈，少用)
// 循环 while-loop

//模板
//4要点
//start + 1 < end  相邻
//start + (end - start)/2	防止溢出
//A[mid] ==, < ,>
//A[start] A[end] ? target
//
//比O(n)更优的时间复杂度几乎只能是O(logn)的二分法
//时间复杂度 log(N)
//二分法
//倍增法（反二分法）

// 排序
// 找最后一个，找第一个

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		mid := start + (end - start)/2  // 防止溢出
		if nums[mid] == target {
			end = mid
		}else if nums[mid] < target {
			start = mid
		}else {
			end = mid
		}
	}

	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}



func main() {
	nums := []int{2, 3, 4, 6, 7, 8, 10, 13}
	idx := search(nums, 4)
	fmt.Print(idx)
}