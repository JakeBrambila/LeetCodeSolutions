package main

import "fmt"

func main() {
	list1 := []int{7, 1, 5, 3, 6, 4}
	list2 := []int{7, 6, 4, 3, 1}
	fmt.Println("Max Profit for list1: ", maxProfit(list1))
	fmt.Println("Max Profit for list2: ", maxProfit(list2))

	fmt.Println()

	fmt.Println("Max Profit for list1: ", maxProfit2(list1))
	fmt.Println("Max Profit for list2: ", maxProfit2(list2))
}

// brute force method
// uses two loops, for every number in the list there is another loop that caluclates the difference
// at every index smaller than the current index and checks that number with the temp variable
// temp then becomes the bigger number
// O(n^2) -> nested loops
func maxProfit(prices []int) int {
	var temp int
	for i := 1; i < len(prices); i++ {
		for j := i - 1; j >= 0; j-- {
			if prices[i]-prices[j] > temp {
				temp = prices[i] - prices[j]
			}
		}
	}
	return temp
}

// efficient solution
// uses sliding pointers to go thorugh the list efficiently
// left points to the smallest current number and right iterates through all the numbers in the list
// O(n) -> only has one loop
func maxProfit2(prices []int) int {
	left := 0
	right := 1
	maxP := 0

	for right < len(prices) {
		if prices[left] < prices[right] {
			if prices[right]-prices[left] > maxP {
				maxP = prices[right] - prices[left]
			}
		} else {
			left = right
		}
		right += 1
	}
	return maxP
}
