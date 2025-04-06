/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dumm := &ListNode{-1, head}
	iter := dumm.Next
	leading := dumm
	for iter != nil {

		curr := iter
		isDuplicate := false
		for curr.Next != nil && curr.Next.Val == iter.Val {
			isDuplicate = true
			curr = curr.Next
		}
		
		if isDuplicate {

			if curr.Next != nil {
				iter = curr.Next
				leading.Next = iter
			} else {
				leading.Next = nil
				break
			}
		} else {
			leading = iter
			iter = iter.Next
		}

	}
	return dumm.Next
}
