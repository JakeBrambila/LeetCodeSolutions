package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var mapNumToIndex = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if index, ok := mapNumToIndex[complement]; ok {
			return []int{i, index}
		}
		mapNumToIndex[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

}
