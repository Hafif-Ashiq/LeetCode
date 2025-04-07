/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == head {
		return head
	}

	arr := make([]*ListNode, 5)

	iter := head

	for iter != nil {
		if slices.Contains(arr, iter) {
			return iter
		}
		arr = append(arr, iter)
		iter = iter.Next
	}

	return nil

}
