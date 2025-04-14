/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dumm := new(ListNode)
	dummEnd := dumm

	iter := head

	for iter != nil {

		it := dumm
		
		for it != nil {
			if it.Next != nil && it.Next.Val > iter.Val {
				next := it.Next
				it.Next = iter
				iter = iter.Next
				it.Next.Next = next
				break
			}
			it = it.Next
		}

		// Is the largest one
		if it == nil {
			dummEnd.Next = iter
			dummEnd = iter
			iter = iter.Next
			dummEnd.Next = nil
		}

	}

	return dumm.Next

}
