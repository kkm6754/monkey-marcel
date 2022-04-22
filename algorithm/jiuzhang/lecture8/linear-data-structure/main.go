package main

import "math"

// 线性数据结构
//	队列
//	栈
//	hash


//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop()—— 删除栈顶的元素。
//top()—— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。





type MinStack struct {
	minStack []int
	valStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minStack: []int{math.MaxInt64}, // 需要分配空间，push需要取值比较大小
		valStack: []int{},
	}
}


func (this *MinStack) Push(val int)  {
	this.valStack = append(this.valStack, val)
	minVal := this.minStack[len(this.minStack) - 1]
	this.minStack = append(this.minStack, min(minVal, val))
}


func (this *MinStack) Pop()  {
	this.valStack = this.valStack[: len(this.valStack) - 1]
	this.minStack = this.minStack[: len(this.minStack) - 1]
}


func (this *MinStack) Top() int {
	return this.valStack[len(this.valStack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack) - 1]
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */