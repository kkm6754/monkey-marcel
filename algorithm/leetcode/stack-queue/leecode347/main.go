package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}

	heap := make([]int, k)
	i := 0
	dealMap := make(map[int]int, 0)
	for key, _ := range m {
		heap[i] = key
		dealMap[key] = 1
		i++
		if i == k {
			break
		}
	}

	for p := (k - 1) / 2; p >= 0; p-- {
		heapf(p, k-1, heap, m)
	}


	for key, val := range m {
		if dealMap[key] == 1 {
			continue
		}
		if val > m[heap[0]] {
			heap[0] = key
			heapf(0, k-1, heap, m)
		}
	}

	return heap
}


func heapf(p, end int, nums []int, m map[int]int) {
	for {
		child := 2 * p + 1
		if child > end {
			break
		}
		if child+1 <= end && m[nums[child+1]] < m[nums[child]] {
			child++
		}
		if m[nums[p]] < m[nums[child]] {
			break
		}
		nums[p], nums[child] = nums[child], nums[p]
		p = child
	}
}

func main() {
	nums := []int{1,1,1,2,2,3}
	res := topKFrequent(nums, 2)
	fmt.Println(res)
}

//////////////

//方法一：小顶堆
func topKFrequent1(nums []int, k int) []int {
	map_num:=map[int]int{}
	//记录每个元素出现的次数
	for _,item:=range nums{
		map_num[item]++
	}
	h:=&IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key,value:=range map_num{
		heap.Push(h,[2]int{key,value})
		if h.Len()>k{
			heap.Pop(h)
		}
	}
	res:=make([]int,k)
	//按顺序返回堆中的元素
	for i:=0;i<k;i++{
		res[k-i-1]=heap.Pop(h).([2]int)[0]
	}
	return res
}

//构建小顶堆
type IHeap [][2]int

func (h IHeap) Len()int {
	return len(h)
}

func (h IHeap) Less (i,j int) bool {
	return h[i][1]<h[j][1]
}

func (h IHeap) Swap(i,j int) {
	h[i],h[j]=h[j],h[i]
}

func (h *IHeap) Push(x interface{}){
	*h=append(*h,x.([2]int))
}
func (h *IHeap) Pop() interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[0:n-1]
	return x
}
