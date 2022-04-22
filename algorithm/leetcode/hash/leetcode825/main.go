package main


import "fmt"

//825. 适龄的朋友


func numFriendRequests(ages []int) int {

	ageMap := map[int]int{}
	for _, age := range ages {
		ageMap[age]++
	}

	res := 0
	for a, aNum := range ageMap {
		for b, bNum := range ageMap {
			if b <= a / 2 + 7 || b > a || (b > 100 && a < 100) {
				continue
			} else if a != b {
				res = res + aNum * bNum
			}else if a == b{
				res = res + aNum * (bNum - 1)
			}
		}
	}

	return res

}

func main() {
	ages := []int{73,106,39,6,26,15,30,100,71,35,46,112,6,60,110}
	fmt.Println(numFriendRequests(ages))
}