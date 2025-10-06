package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Next: head}
    current := dummy

    for current.Next != nil {
        if current.Next.Val == val{
            current.Next = current.Next.Next
        }else{
            current = current.Next
        }

    }

    return dummy.Next
}

func deleteNode(node *ListNode) {

   node.Val = node.Next.Val
   node.Next = node.Next.Next


}