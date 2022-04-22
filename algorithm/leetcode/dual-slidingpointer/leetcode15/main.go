package main

import (
	"fmt"
	"sort"
)

// 去重处理
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	sort.Ints(nums)

	for i := 0; i < length - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}

		l, r := i+1, length-1
		for l < r {
			if nums[i] + nums[l] + nums[r] == 0 {
				re := make([]int, 3)
				re[0], re[1], re[2] = nums[i], nums[l], nums[r]
				res = append(res, re)
				for l < r && nums[l] == re[1] { // 去重
					l++
				}
				for l < r && nums[r] == re[2] { // 去重
					r--
				}
			} else if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else {
				l++
			}
		}

	}

	return res
}

func main() {
	nums := []int{-1,0,1,2,-1,-4}

	fmt.Println(threeSum(nums))
}