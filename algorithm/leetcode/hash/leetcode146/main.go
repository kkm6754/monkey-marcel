package main

import "fmt"

type LRUCache struct {
	size int
	capacity int
	cache map[int]*ListNode
	head, tail *ListNode
}

type ListNode struct {
	key, val int
	prev, next *ListNode
}


func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		size: 0,
		capacity: capacity,
		cache: map[int]*ListNode{},
		head: &ListNode{},
		tail: &ListNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok { // 不存在
		node := &ListNode{
			key: key,
			val: value,
		}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			deleteNode := this.tail.prev
			this.removeNode(deleteNode)
			delete(this.cache, key)
			this.size--
		}

	} else { // 存在
		node := this.cache[key]
		node.val = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) removeNode(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 放队列头
func (this *LRUCache)addToHead(node *ListNode)  {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

// 移动到对头
func (this *LRUCache)moveToHead(node *ListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func main() {
	lru := Constructor(3)
	lru.Put(1, 2)
	lru.Put(2, 3)
	lru.Put(3, 4)
	fmt.Println(lru.Get(1))
}


//----
//
//type LRUCache struct {
//	size int
//	capacity int
//	cache map[int]*DLinkedNode
//	head, tail *DLinkedNode
//}
//
//type DLinkedNode struct {
//	key, value int
//	prev, next *DLinkedNode
//}
//
//func initDLinkedNode(key, value int) *DLinkedNode {
//	return &DLinkedNode{
//		key: key,
//		value: value,
//	}
//}
//
//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		cache: map[int]*DLinkedNode{},
//		head: initDLinkedNode(0, 0),
//		tail: initDLinkedNode(0, 0),
//		capacity: capacity,
//	}
//	l.head.next = l.tail
//	l.tail.prev = l.head
//	return l
//}
//
//func (this *LRUCache) Get(key int) int {
//	if _, ok := this.cache[key]; !ok {
//		return -1
//	}
//	node := this.cache[key]
//	this.moveToHead(node)
//	return node.value
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//	if _, ok := this.cache[key]; !ok {
//		node := initDLinkedNode(key, value)
//		this.cache[key] = node
//		this.addToHead(node)
//		this.size++
//		if this.size > this.capacity {
//			removed := this.removeTail()
//			delete(this.cache, removed.key)
//			this.size--
//		}
//	} else {
//		node := this.cache[key]
//		node.value = value
//		this.moveToHead(node)
//	}
//}
//
//func (this *LRUCache) addToHead(node *DLinkedNode) {
//	node.prev = this.head
//	node.next = this.head.next
//	this.head.next.prev = node
//	this.head.next = node
//}
//
//func (this *LRUCache) removeNode(node *DLinkedNode) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//}
//
//func (this *LRUCache) moveToHead(node *DLinkedNode) {
//	this.removeNode(node)
//	this.addToHead(node)
//}
//
//func (this *LRUCache) removeTail() *DLinkedNode {
//	node := this.tail.prev
//	this.removeNode(node)
//	return node
//}

