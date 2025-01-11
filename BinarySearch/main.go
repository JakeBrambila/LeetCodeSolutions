package main

import "fmt"

func main() {
	n1 := []int{ 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(search(n1, 5))

}

// runs a while loop and keeps dividing the array in 2 to find the number in the sorted array
// O(log(n)) because it keep dividing by 2 until it finds the answer
// uses the two pointer approach
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
