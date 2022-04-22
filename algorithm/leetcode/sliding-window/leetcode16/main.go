package leetcode16

import "sort"

//16. 最接近的三数之和


// 排序
// 定下一个元素，剩下2个数，在剩下的元素从前后夹筛选



func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)

	res := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums) - 2; i++ {
		n1 := nums[i]
		l , r := i+1 , len(nums) - 1
		for l < r {
			n2 , n3 := nums[l] , nums[r]
			if n1 + n2 + n3 < target { // 前后夹
				l++
			} else {
				r--
			}
			temp := n1 + n2 + n3
			if abs(temp - target) < abs(res - target) {
				res = temp
			}
		}
	}

	return res

}


// 绝对值
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
