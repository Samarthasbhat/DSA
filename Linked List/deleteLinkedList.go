package main

import "fmt"

// Definition for singly-linked list
type ListNode struct {
    Val  int
    Next *ListNode
}

// Remove all nodes with a specific value
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Next: head}
    current := dummy

    for current.Next != nil {
        if current.Next.Val == val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }

    return dummy.Next
}

// Delete a node when only that node is given
func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
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
    // Example 1: removeElements
    fmt.Println("Remove elements example:")
    head1 := buildList([]int{1, 2, 6, 3, 4, 5, 6})
    fmt.Print("Original list: ")
    printList(head1)

    head1 = removeElements(head1, 6)
    fmt.Print("After removing 6: ")
    printList(head1)

    // Example 2: deleteNode
    fmt.Println("\nDelete given node example:")
    head2 := buildList([]int{4, 5, 1, 9})
    fmt.Print("Original list: ")
    printList(head2)

    nodeToDelete := head2.Next // delete node with value 5
    deleteNode(nodeToDelete)
    fmt.Print("After deleting node 5: ")
    printList(head2)
}
