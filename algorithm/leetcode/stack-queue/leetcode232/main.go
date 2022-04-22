package main
//
//type MyQueue struct {
//	stackIn []int
//	stackOut []int
//	size int
//}
//
//
//func Constructor() MyQueue {
//	return MyQueue{
//		make([]int, 0),
//		make([]int, 0),
//		0,
//	}
//}
//
//
//func (this *MyQueue) Push(x int)  {
//	if len(this.stackOut) > 0 {
//		for i := this.size - 1; i >= 0; i-- {
//			this.stackIn = append(this.stackIn, this.stackOut[i])
//		}
//		this.stackOut = []int{}
//	}
//	this.stackIn = append(this.stackIn, x)
//	this.size++
//}
//
//
//func (this *MyQueue) Pop() int {
//	res := this.Peek()
//	this.stackOut = this.stackOut[: this.size - 1]
//	this.size--
//	return res
//}
//
//
//func (this *MyQueue) Peek() int {
//	if this.size == 0 {
//		return 0
//	}
//	if len(this.stackIn) > 0 {
//		for i := this.size - 1; i >= 0; i-- {
//			this.stackOut = append(this.stackOut, this.stackIn[i])
//		}
//		this.stackIn = []int{}
//	}
//	res := this.stackOut[this.size - 1]
//	return res
//}
//
//
//func (this *MyQueue) Empty() bool {
//	return this.size == 0
//}
//
