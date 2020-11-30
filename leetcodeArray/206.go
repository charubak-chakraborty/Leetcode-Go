package leetcodeArray

// Reverse a singly linked list.

// Example:

// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
// Follow up:

// A linked list can be reversed either iteratively or recursively. Could you implement both?

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative
// func ReverseList(head *ListNode) *ListNode {
// 	var pre, cur *ListNode = nil, head
// 	for cur != nil {
// 		next := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = next
// 	}
// 	return pre
// }

// recursive
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	r := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return r
}

// 2->4->3
// pre = nil cur = 2

// next = cur.next = 4
// cur.next = 2.next = pre = nil
// pre =  cur = 2
// cur = next = 4

// 4->nil
