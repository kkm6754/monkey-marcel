package leetcode128

//128. 最长连续序列


func longestConsecutive(nums []int) int {

	// 所有数的集合
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}

	res := 0

	for num := range numMap {
		if !numMap[num - 1] { // num-1 不存在，表示num可以作为序列的起始，这里可以排除连续数后面数的计算
			cur := num
			curLen := 1
			for numMap[cur + 1] { // 以num开始，往后查看后续最多连续序列
				cur++
				curLen++
			}
			if res < curLen {
				res = curLen
			}
		}
	}

	return res
}