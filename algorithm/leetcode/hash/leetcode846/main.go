package main

import (
	"fmt"
	"sort"
)

//846. 一手顺子

func isNStraightHand(hand []int, groupSize int) bool {
	fmt.Println(hand)
	numMap := map[int]int{}
	for _, v := range hand {
		numMap[v]++
	}

	sort.Ints(hand)
	for _, k := range hand {
		if numMap[k] > 0 {
			for i := 0; i < groupSize; i++ {
				if numMap[k + i] == 0 {
					return false
				}else {
					numMap[k + i]--
				}
			}
		}
	}

	for _, v := range numMap {
		if v > 0{
			return false
		}
	}

	return true

}

func main() {
	hand := []int{8, 10, 12}
	fmt.Println(isNStraightHand(hand, 3))
}