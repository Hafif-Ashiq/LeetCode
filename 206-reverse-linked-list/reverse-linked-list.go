/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	last := head

	for last.Next != nil {
		last = last.Next
	}

	anchorL := last

	for {
		iter := head

		for iter.Next != anchorL {
			iter = iter.Next
		}
		anchorL.Next = iter
		anchorL = iter

		if anchorL == head {
			break
		}
	}
	head.Next = nil
	return last
}