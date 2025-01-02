package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := createList([]int{1, 3, 5})
	list2 := createList([]int{2, 4, 6})

	mergedList := mergeTwoLists(list1, list2)

	printList(mergedList)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//go defaults 0 for Val and nil for dummy.Next
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		//move tail to the new node
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	} else if list2 == nil {
		tail.Next = list1
	}

	return dummy.Next
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
