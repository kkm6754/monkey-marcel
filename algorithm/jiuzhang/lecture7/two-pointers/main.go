package main

// 双指针
// 2数和、3数和、4数和、k数和
// 排序



// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// hashMap 遍历数组，将差值作为key存入hashmap， 继续遍历，通过后续节点的差值是否存在hashmap中，得到结果
// 排序 + 双指针 复制新数组出来，排序后，双指针找到2个数，再从原数组中找到下标



func twoSum(nums []int, target int) []int {
	ret := make([]int, 0)
	if len(nums) < 2 {
		return ret
	}

	vIdxHash := map[int]int{}
	for idx, v := range nums {
		if  idx2, ok := vIdxHash[target - v]; ok {
			return []int{idx2,idx}
		}
		vIdxHash[v] = idx
	}

	return ret
}






