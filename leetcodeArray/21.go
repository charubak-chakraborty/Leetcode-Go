package leetcodeArray

// Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

// Example:

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var output = &ListNode{}
	var ptr = output

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr.Next = l1
			l1 = l1.Next
			ptr = ptr.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
			ptr = ptr.Next
		}
	}
	if l1 == nil {
		ptr.Next = l2
	}
	if l2 == nil {
		ptr.Next = l1
	}
	return output.Next
}
