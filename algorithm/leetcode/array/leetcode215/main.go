package main


import "fmt"

func findKthLargest(nums []int, k int) int {

	heapSort(nums)

	fmt.Println(nums)
	return nums[k-1]
}


func heapSort(nums []int) {
	end := len(nums) - 1

	// 堆化
	for p := end / 2; p >= 0; p-- {
		heapf(p, end, nums)
	}
	fmt.Println(nums)

	// 不断的取最小值放到最后，最终返回一个由大到小排序
	for i := end; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		end--
		heapf(0, end, nums)
	}


}

func heapf(p, end int, nums []int) {
	for {
		child := 2 * p + 1 //左子节点
		if child > end {
			break
		}
		// 右子节点存在，且比左子节点小
		if child + 1 <= end && nums[child + 1] < nums[child] {
			child++
		}
		// 父节点比左右子节点都小
		if nums[p] < nums[child]{
			break
		}
		nums[p], nums[child] = nums[child], nums[p]
		p = child
	}
}




func main() {
	num := []int{3,2,3,1,2,4,5,5,6}
	//ky := findKthLargest(num, 4)

	//SelectSort(num)
	//bubbleSort(num)
	//insertionSort(num)
	quickSort(num)


	fmt.Println(num)
}


func SelectSort(nums []int){
	l := len(nums)
	if l < 1 {
		return
	}

	for i := 0; i < l - 1; i++ {
		maxIdx := i
		for j := i; j < l; j++ {
			if nums[maxIdx] < nums[j]{
				maxIdx = j
			}
		}
		if maxIdx != i {
			nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
		}
	}

}

func bubbleSort(nums []int) {
	l := len(nums)
	if l < 1 {
		return
	}

	for i := l - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] < nums[j + 1] {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}
}


func insertionSort(arr []int) []int {
	for i := range arr { // 从第一个元素开始
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] < current { // 第二个元素开始比较，找到位置前，平移数据
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current // 插入数据
	}
	return arr
}


func mergeSort(s []int) []int{
	len := len(s)
	if len == 1 {
		return s //最后切割只剩下一个元素
	}

	// 递归
	m := len/2
	leftS := mergeSort(s[:m])
	rightS := mergeSort(s[m:])

	return merge(leftS, rightS)
}

//把两个有序切片合并成一个有序切片
func merge(l []int, r []int) []int{
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)

	lIndex,rIndex := 0,0 //两个切片的下标，插入一个数据，下标加一
	for lIndex<lLen && rIndex<rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		}else{
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}


func quickSort(nums []int){
	quickSort0(nums, 0 , len(nums) - 1)
}

func quickSort0(nums []int, left, right int){
	if left < right {
		partitionIdx := partition(nums, left, right)

		quickSort0(nums, left, partitionIdx - 1)
		quickSort0(nums, partitionIdx + 1, right)
	}
}

func partition(nums []int, left, right int) int{
	pivot := left
	idx := pivot + 1
	for i := idx; i <= right; i++ {
		if nums[i] > nums[pivot] {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}
	nums[pivot], nums[idx - 1] = nums[idx - 1], nums[pivot]
	return idx - 1
}

