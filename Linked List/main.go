package main

import "fmt"

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates removes all duplicate numbers (LeetCode #82)
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // create dummy node
	prev := dummy                  // last distinct node
	current := head

	for current != nil {
		// Move current until the end of duplicates group
		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next
		}

		// If duplicates were found (prev.Next != current)
        // skip them all
		if prev.Next != current {
			prev.Next = current.Next
		} else {
			prev = prev.Next
		}

		current = current.Next
	}

	return dummy.Next
}

// Helper: Build linked list from slice
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// Helper: Print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example 1
	input := []int{1, 2, 3, 3, 4, 4, 5}
	head := buildList(input)

	fmt.Println("Original list:")
	printList(head)

	head = deleteDuplicates(head)

	fmt.Println("After removing duplicates:")
	printList(head)

	// Example 2
	input2 := []int{1, 1, 1, 2, 3}
	head2 := buildList(input2)

	fmt.Println("\nOriginal list:")
	printList(head2)

	head2 = deleteDuplicates(head2)

	fmt.Println("After removing duplicates:")
	printList(head2)
}
