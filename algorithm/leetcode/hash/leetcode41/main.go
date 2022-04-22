package leetcode41


//41. 缺失的第一个正数

func firstMissingPositive(nums []int) int {

	l := len(nums)

	for i := 0; i < l; i++ {
		for nums[i] > 0 && nums[i] <= l && nums[i] != nums[nums[i] - 1] { // nums[i] != nums[nums[i] - 1] 需要校验是否置换
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}
	}

	for i, val := range nums {
		if i + 1 != val {
			return i + 1
		}
	}

	return l + 1
}