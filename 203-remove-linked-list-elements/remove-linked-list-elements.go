/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil || (head.Next == nil && head.Val == val) {
		return nil
	}
	iter := head
	var leading *ListNode

	for iter != nil {
		if iter.Val == val {
			if iter == head {
				head = iter.Next
				iter = iter.Next
			} else {
				leading.Next = iter.Next
				iter = iter.Next
			}

		} else {

			leading = iter
			iter = iter.Next
		}

	}

	return head
}