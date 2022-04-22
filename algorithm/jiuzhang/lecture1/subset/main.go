package main


import "fmt"

// subsets 全子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集

// 回溯法
// 把一个数放入集合，然后又将其取出，使集合状态回到上一层状态


func subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	if len(nums) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)
	path := make([]int, 0, len(nums))
	dfs(path, nums, 0, &result)
	
	return result
}


func dfs(path []int, nums []int, cur int, ret *[][]int) {
	tmp := make([]int, len(path)) // tmp := make([]int, 0, len(path)) 最后结果错误 [[] [] [] [] [] [] [] []] 
	copy(tmp, path) 			  // copy 复制的数量由dest分片的size决定，不会自动扩容
	//fmt.Printf("tmp: %v\n", tmp)
	*ret = append(*ret, tmp)
	// fmt.Printf("ret: %v\n", *ret)

	for i := cur; i < len(nums); i++ {
		path = append(path, nums[i])
		fmt.Printf("path: %v\n", path)
		dfs(path, nums, i + 1, ret)
		path = path[:len(path)-1]  // 回溯
	} 

}



func main(){
	nums := []int{1, 2, 3}
	ret := subsets(nums)
	fmt.Println(ret)
}