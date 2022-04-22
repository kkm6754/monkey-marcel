package leetcode88

// 88. 合并两个有序数组


func merge(nums1 []int, m int, nums2 []int, n int)  {

	i := m + n - 1
	p1 := m - 1
	p2 := n - 1


	// or，2者其中之一有数据就继续执行
	for p1 >= 0 || p2 >= 0 {
		if p1 == -1 {
			nums1[i] = nums2[p2]
			p2--
		}else if p2 == -1 {
			nums1[i] = nums1[p1]
			p1--
		}else if nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		}else {
			nums1[i] = nums2[p2]
			p2--
		}
		i--
	}

}