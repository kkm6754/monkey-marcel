package main


// 排序数组
// 合并2个排序的数组，合并k个排序的数组
// 2个排序数组的中位数



// 合并有序数组
// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n

// 避免整体拖动，先放大的数到最后位置


// 扩展 Quick Select - 无序数组的第K大数，O(n)


func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}

	for i := len(nums1); i > 0; i-- {
		if m > 0 && n > 0 {
			if nums1[m - 1] > nums2[n-1] {
				nums1[i - 1] = nums1[m - 1]
				m--
			} else {
				nums1[i - 1] = nums2[n - 1]
				n--
			}
		}else if m > 0{
			nums1[i - 1] = nums1[m - 1]
			m--
		}else if n > 0 {
			nums1[i - 1] = nums2[n - 1]
			n--
		}
	}


	return
}


