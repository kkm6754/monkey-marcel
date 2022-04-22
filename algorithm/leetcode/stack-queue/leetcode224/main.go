package leetcode224


func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	queue := make([]int, 0)

	for i := 0; i < k; i++ {
		queue = append(queue, nums[i])
		if len(queue) > 1 && queue[0] < nums[i] {
			queue = queue[len(queue) - 1:]
		}
	}
	res = append(res, queue[0])

	for i := k; i < len(nums); i++ {
		// pop
		if queue[0] == nums[i - k] {
			queue = queue[1:]
		}

		// push
		for len(queue) > 0 && queue[len(queue) - 1] < nums[i] {
			queue = queue[: len(queue) - 1]
		}
		queue = append(queue, nums[i])


		res = append(res, queue[0])
	}

	return res

}




/////////////////////////

// 封装单调队列的方式解题

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow2(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}



/////////////////////////


func maxSlidingWindow1(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	max := 0
	res := make([]int, 0)
	queue := make([]int, 0)

	for i := 0; i < k; i++ {
		queue = append(queue, nums[i])
		if nums[i] > max {
			max = nums[i]
		}
	}
	res = append(res, max)

	for i := k; i < len(nums); i++ {
		queue = queue[1:]
		queue = append(queue, nums[i])

		if nums[i] > max {
			max = nums[i]
		}
		res = append(res, max)
	}

	return res
}
