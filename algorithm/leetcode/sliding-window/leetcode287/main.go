package leetcode287


//287. 寻找重复数


// 环检测算法
//

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	// 寻找环内的相遇点
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] { }

	// 同部长，寻找环的起始点（重复的数字）
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}


